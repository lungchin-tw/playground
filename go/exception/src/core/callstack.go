package core

import (
	"fmt"
	"runtime"
)

func getFrame(skip int) runtime.Frame {
	pc := make([]uintptr, 1)
	num := runtime.Callers(skip, pc)
	if num <= 0 {
		return runtime.Frame{}
	}

	frame, _ := runtime.CallersFrames(pc).Next()
	return frame
}

func CurFuncName() string {
	return getFrame(3).Function
}

func CurFuncDesc() string {
	frame := getFrame(3)
	return fmt.Sprintf("%v:%v", frame.Function, frame.Line)
}
