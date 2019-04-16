package main

import "time"

import "fmt"

type TickerInterface interface {
	Signal() chan uint32
	Quit() chan bool
}

type TickerHandler struct {
	step     uint32
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
	if c.StepNumber() > c.step {
		fmt.Printf("timeToNext : %+v\n", c.TimeLaps/2)
		return c.TimeLaps / 2
	}
	tmp := c.refTime.Add(time.Duration(c.step+1) * c.TimeLaps).Sub(time.Now())
	if tmp > 0 {
		fmt.Printf("timeToNext : %+v\n", tmp)
		return tmp
	}
	fmt.Printf("timeToNext : %+v\n", time.Millisecond)
	return time.Duration(1 * time.Millisecond)
}

func (c *TickerHandler) StepNumber() uint32 {
	return uint32(time.Now().Sub(c.refTime) / c.TimeLaps)
}

func (c *TickerHandler) Step() uint32 {
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
			stateRunning(c, i)
		case <-c.quit:
			return
		}
	}
	close(c.start)
	close(c.pause)
	close(c.quit)
	i.Quit() <- true
}

func stateRunning(c *TickerHandler, i TickerInterface) {
	for {
		fmt.Println(" Loop Running State.")
		select {
		case <-c.pause:
			fmt.Println(" Pause while running")
			c.pauseAt = time.Now()
			return
		case <-time.After(c.timeToNext()):
			fmt.Println(" commit Signal.")
			commitSignal(c, i)
		case <-c.quit:
			fmt.Println(" Quit while running")
			return
		}
		fmt.Println("end loop Running State.")
	}
	fmt.Println("quit Running State.")
}

func commitSignal(c *TickerHandler, i TickerInterface) {
	select {
	case i.Signal() <- c.step:
		c.step += 1
	default:
		// no commitment
	}
}
