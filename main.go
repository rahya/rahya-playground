package main

import "fmt"
import "time"

var JS go2js

type worker struct {
	t *TickerHandler
}

func (w *worker) Action(h WorkerHandlerInterface) {
	JS.call(JS_TEST, h.Step(), w.t.Step())
}

func (w *worker) Close(h WorkerHandlerInterface) {
}

func main() {
	JS.Init()

	fmt.Println("Let's go, Go !")

	tHandler := &TickerHandler{TimeLaps: 50 * time.Millisecond}
	tHandler.Init()

	wHandler := &WorkerHandler{}
	wHandler.Init()
	work := &worker{t: tHandler}

	go ThreadWorker(wHandler, work)
	go Thread(tHandler, wHandler)

	tHandler.Start()
	time.Sleep(100 * time.Second)
	tHandler.Pause()
	fmt.Println("Finish, Go !")

	var ch chan bool
	ch <- true

}
