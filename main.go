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

func stateRunning() {
	var td time.Duration
	var timeLaps time.Duration = 50 * time.Millisecond
	var refTime time.Time = time.Now()
	var step int = 0
	for {
		td = refTime.Add(time.Duration(step+1) * timeLaps).Sub(time.Now())
		fmt.Printf("td: %+v\n", td)
		time.Sleep(td)
		{
			//i.Signal() <- c.step
			JS.call(JS_TEST, step, 0)
			step += 1
		}
	}
}

func main() {
	JS.Init()

	fmt.Println("Let's go, Go !")

	//tHandler := &TickerHandler{TimeLaps: 50 * time.Millisecond}
	//tHandler.Init()

	//wHandler := &WorkerHandler{}
	///wHandler.Init()
	//work := &worker{t: tHandler}

	//go ThreadWorker(wHandler, work)
	//tHandler.refTime = time.Now()
	//tHandler.pauseAt = time.Now()
	go stateRunning() //wHandler)

	//tHandler.Start()
	//time.Sleep(100 * time.Second)
	//tHandler.Pause()
	fmt.Println("Finish, Go !")
	// */

	var ch chan bool
	ch <- true

}
