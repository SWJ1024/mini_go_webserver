package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopl.io/mini-videoserver/scheduler/dbops"
)


func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		SendResponse(w, 400, "video is should not be empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		SendResponse(w, 500, "Internal Server error")
		return
	}
	SendResponse(w, 200, "")
	return
}