package main

//import "fmt"
import "time"

func debug(str string) {
	//fmt.Println(str)

}

type WorkerHandler struct {
	step   int
	signal chan int
	quit   chan bool
}

func (ossd *WorkerHandler) Init() {
	ossd.signal = make(chan int)
	ossd.quit = make(chan bool)
	ossd.step = 0
}

func (ossd *WorkerHandler) StepSet(s int) {
	ossd.step = s
}

func (ossd *WorkerHandler) Step() int {
	return ossd.step
}

func (ossd *WorkerHandler) Close() {
	close(ossd.signal)
	close(ossd.quit)
}

func (ossd *WorkerHandler) Signal() chan int {
	return ossd.signal
}

func (ossd *WorkerHandler) Quit() chan bool {
	return ossd.quit
}

//*
func (ossd *WorkerHandler) Action(h WorkerHandlerInterface, bi WorkerInterface) {
	bi.Action(h)
}

//*/
type WorkerHandlerInterface interface {
	Action(h WorkerHandlerInterface, bi WorkerInterface)
	Init()
	StepSet(s int)
	Step() int
	Close()
	Signal() chan int
	Quit() chan bool
}

type WorkerInterface interface {
	Action(h WorkerHandlerInterface)
	Close(h WorkerHandlerInterface)
}

func ActionWorker(h WorkerHandlerInterface, bi WorkerInterface) {
	bi.Action(h)
}

func ThreadWorker(h WorkerHandlerInterface, bi WorkerInterface) {
	for {
		debug("loop ThreadWorker")
		select {
		case step := <-h.Signal():
			debug("loop step ThreadWorker")
			h.StepSet(step)
			h.Action(h, bi)
			debug("loop step end ThreadWorker")
		case <-h.Quit():
			debug("loop quit ThreadWorker")
			return
		case <-time.After(time.Second):
			debug("Still alive")
		}

		debug("loop end ThreadWorker")
	}
	h.Close()
	bi.Close(h)
	debug("quit ThreadWorker")
}
