package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VideoDir + vid
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	defer video.Close()
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	//log.Printf("entered the streamhandler")
	//targetUrl := "https://go-lessions.oss-cn-beijing.aliyuncs.com/videos" + p.ByName("vid-id")
	//http.Redirect(w, r, targetUrl, 301)
}


func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "file is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	filename := p.ByName("vid-id")
	err = ioutil.WriteFile(VideoDir+ filename, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	/*
	ossfn := "videos/" + filename
	path := "./videos/" + filename
	bn := "go-lessions.oss-cn-beijing.aliyuncs.com"
	if !UploadToOss(ossfn, path, bn)  {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	os.Remove(path)
*/
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully!")
}


func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}