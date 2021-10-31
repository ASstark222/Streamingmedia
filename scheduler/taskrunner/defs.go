package taskrunner

type controlChan chan string //control channel

type dataChan chan interface{} //data channel

type fn func(dc dataChan) error  //dispatcher和executer

//三个类型的消息
const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

)

const VIDEO_PATH = "./videos/"