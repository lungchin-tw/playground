package env_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"testing"
)

func TestEnv(t *testing.T) {
	dir, err := os.Getwd()
	fmt.Printf("os.Getwd(): %v, Error: %v\n ", dir, err)
	fmt.Println("os.Args: ", os.Args)
	fmt.Println("runtime.GOOS: ", runtime.GOOS)
	fmt.Println("runtime.GOARCH: ", runtime.GOARCH)
	fmt.Println("runtime.GOROOT: ", runtime.GOROOT())
	fmt.Println("runtime.NumCPU: ", runtime.NumCPU())
	fmt.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	fmt.Println("runtime.GOMAXPROCS(-1): ", runtime.GOMAXPROCS(-1))
	fmt.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	fmt.Println("runtime.GOMAXPROCS(2): ", runtime.GOMAXPROCS(2))
	fmt.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	fmt.Println("runtime.NumCgoCall: ", runtime.NumCgoCall())
	fmt.Println("runtime.NumGoroutine: ", runtime.NumGoroutine())
	fmt.Println("runtime.Version: ", runtime.Version())
	assert.True(true)
}
