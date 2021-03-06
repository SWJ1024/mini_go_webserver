package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"

	"gopl.io/mini-videoserver/scheduler/dbops"
)


func deleteVideo(vid string) error {
	err := os.Remove(VideoPath + vid)
	if err != nil && !os.IsNotExist(err){
		log.Printf("Delete video error: %v", err)
		return err
	}
	return nil
}


func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("VideoClearDispatcher error : %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
		}
	return nil
}


func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}

	forloop:
		for {
			select {
			case vid := <-dc:
				go func(id interface{}) {
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
			default:
				break forloop
			}
		}

	var err error
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}