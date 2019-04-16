package main

import (
	"syscall/js"
)

const (
	JS_START             = iota
	JS_CONSOLE_LOG       = iota
	JS_TEST              = iota
	CALL_NUM       uint8 = iota
)

type go2js struct {
	calls [CALL_NUM]js.Value
}

func (ji *go2js) Init() {
	ji.registerCalls()

}

func (ji *go2js) registerCalls() {
	ji.calls[JS_START] = js.Global().Get("start")
	ji.calls[JS_TEST] = js.Global().Get("wasm_test")
	ji.calls[JS_CONSOLE_LOG] = js.Global().Get("console_log")

	//fmt.Printf("Function %+T\n", pp)
}

func (ji *go2js) call(pos uint8, args ...interface{}) {
	ji.calls[pos].Invoke(args...)
}

func (ji *go2js) ConsoleLog(s string) {
	ji.call(JS_CONSOLE_LOG, s)
}
