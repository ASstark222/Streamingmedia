package taskrunner

/*
在runner里跑常驻的任务
 */

type Runner struct {
	Controller controlChan
	Error controlChan //Error channel返回CLOSE信息
	Data dataChan
	dataSize int
	longlived bool //是否长期保留资源
	Dispathcher fn
	Executor fn
}


//初始化一个Runner
func NewRunner(size int,longlived bool,d,e fn) *Runner {
	return &Runner{Controller: make(chan string,1), //整个过程是非阻塞的
					Error: make(chan string,1),
					Data: make(chan interface{},size),
					dataSize: size,
					longlived: longlived,
					Dispathcher: d,
					Executor: e,
		}
}

func (r *Runner) startDispatch()  {
	defer func() {
		if !r.longlived{
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
			case c :=<-r.Controller:
				if c == READY_TO_DISPATCH{
					err :=r.Dispathcher(r.Data)
					if err != nil{
						r.Error <-CLOSE
					}else{ //分配完毕，告诉executor执行任务
						r.Controller<-READY_TO_EXECUTE
					}
				}
				if c == READY_TO_EXECUTE {
					err :=r.Executor(r.Data)
					if err !=nil{
						r.Error<-CLOSE
					}else { //执行正确且完毕，可以分配下一个任务
						r.Controller<-READY_TO_DISPATCH
					}
				}
			case e :=<-r.Controller:
				if e == CLOSE{
					return
				}
		default:
		}
	}
}

func (r *Runner) startAll()  {
	r.Controller<-READY_TO_DISPATCH
	r.startDispatch()
}