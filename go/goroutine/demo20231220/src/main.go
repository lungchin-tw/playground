package main

import (
	"fmt"
	"net/http"
	"os"
	"playground/goroutine/demo20231220/util"

	"github.com/gin-gonic/gin"
)

var JobQueue chan string

func init() {
	JobQueue = make(chan string, 100)
}

func main() {
	path, _ := os.Getwd()
	fmt.Println(util.CurFuncDesc(), path)

	NewDispatcher(JobQueue, 32).Run(4)

	router := gin.Default()
	router.Handle(http.MethodPost, "/submit", submit)
	router.Run(":8080")
}

func submit(ctx *gin.Context) {
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.String(http.StatusBadRequest, "%v", "Bad Request")
		return
	}

	JobQueue <- ctx.PostForm("message")
	ctx.String(http.StatusOK, "%v", util.CurFuncDesc())
}
