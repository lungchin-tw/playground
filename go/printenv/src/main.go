package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	var val interface{}
	var err error

	val, err = os.Getwd()
	log.Println("os.Getwd():", val, err)

	log.Println("os.Args:", os.Args)

	val, err = os.Executable()
	log.Println("os.Executable:", val, err)

	val, err = filepath.Abs("./")
	log.Println("filepath.Abs(\"./\"):", val, err)

	_, val, _, ok := runtime.Caller(0)
	log.Println("runtime.Caller(0):", val, ok)

	log.Println("runtime.GOOS: ", runtime.GOOS)
	log.Println("runtime.GOARCH: ", runtime.GOARCH)
	log.Println("runtime.GOROOT: ", runtime.GOROOT())
	log.Println("runtime.NumCPU: ", runtime.NumCPU())
	log.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	log.Println("runtime.GOMAXPROCS(-1): ", runtime.GOMAXPROCS(-1))
	log.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	log.Println("runtime.GOMAXPROCS(2): ", runtime.GOMAXPROCS(2))
	log.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	log.Println("runtime.NumCgoCall: ", runtime.NumCgoCall())
	log.Println("runtime.NumGoroutine: ", runtime.NumGoroutine())
	log.Println("runtime.Version: ", runtime.Version())
}
