package taskrunner

import (
	"Streamingmedia/scheduler/dbops"
	"Streamingmedia/scheduler/ossops"
	"errors"
	"log"
	"sync"
)


//删除Video
func deleteVideo(vid string) error {

	ossfn := "videos/"+vid
	bn :="govideos"

	ok := ossops.DeleteObject(ossfn,bn)
	if !ok{
		log.Printf("deleeting video error,oss fail:%s",ok)
		return errors.New("Deleting video error")
	}

	return nil
}


/*
定期分配删除Video 任务
 */
func VideoClearDispatcher(dc dataChan) error  {
	res,err := dbops.ReadVideoDeletionRecord(3)
	if err !=nil{
		log.Printf("video clear dispatcher error:%v",err)
		return err
	}

	if len(res) == 0{
		return errors.New("All tasks finished")
	}

	//将id一个个送入datachannel
	for _,id := range res{
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
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

	errMap.Range(func(k, v interface{}) bool {
		err = k.(error)
		if err != nil{
			return false
		}
		return true
	})

	return err
}