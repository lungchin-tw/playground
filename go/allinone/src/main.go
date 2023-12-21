package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"playground/allinone/api"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	panicerr := make(chan error)
	path, _ := os.Getwd()
	fmt.Println(path)

	http.HandleFunc("/health", api.HandleHealth)
	http.HandleFunc("/add/match", api.AddSinglePersonAndMatch)
	http.HandleFunc("/remove/user", api.RemoveSinglePerson)
	http.HandleFunc("/query/user", api.QuerySinglePerson)

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panicerr <- err
		}
	}()

	fmt.Println("Web Server Startup.")
	panic(<-panicerr)
}
