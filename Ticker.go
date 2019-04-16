package main

import "time"

//import "fmt"

type TickerInterface interface {
	Signal() chan int
	Quit() chan bool
}

type TickerHandler struct {
	step     int
	refTime  time.Time
	TimeLaps time.Duration
	pauseAt  time.Time
	start    chan bool
	pause    chan bool
	quit     chan bool
}

func (c *TickerHandler) Start() {
	c.start <- true
}

func (c *TickerHandler) Pause() {
	c.pause <- true
}

func (c *TickerHandler) Quit() {
	c.quit <- true
}

func (c *TickerHandler) Init() {
	c.start = make(chan bool)
	c.pause = make(chan bool)
	c.quit = make(chan bool)
}

func (c *TickerHandler) timeToNext() time.Duration {
	tmp := c.refTime.Add(time.Duration(c.step+1) * c.TimeLaps).Sub(time.Now())
	return tmp
}

func (c *TickerHandler) StepNumber() int {
	return int(time.Now().Sub(c.refTime) / c.TimeLaps)
}

func (c *TickerHandler) Step() int {
	return c.step
}

func Thread(c *TickerHandler, i TickerInterface) {
	c.refTime = time.Now()
	stateWait(c, i)
}

func stateWait(c *TickerHandler, i TickerInterface) {
	c.pauseAt = time.Now()
	for {
		select {
		case <-c.pause:
			// Ignore and should never be call
		case <-c.start:
			c.refTime = c.refTime.Add(time.Now().Sub(c.pauseAt))
			//stateRunning(c, i)
		case <-c.quit:
			return
		}
	}
	close(c.start)
	close(c.pause)
	close(c.quit)
	//i.Quit() <- true
}

/*
func stateRunning(c *TickerHandler) {
	var td time.Duration
	for {
		td = c.timeToNext()
		fmt.Printf("td: %+v\n", td)
		time.Sleep(td)
		{
			//i.Signal() <- c.step
			JS.call(JS_TEST, c.Step(), 0)
			c.step += 1
		}
	}
}
*/

func commitSignal(c *TickerHandler, i TickerInterface) {
	/*
		select {
		case i.Signal() <- c.step:
			c.step += 1
		default:
			// no commitment
		}
	*/
}
