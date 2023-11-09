package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"playground/grpc/core"
	"playground/grpc/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {
	address := flag.String("address", "", "the server address")
	flag.Parse()
	log.Println("Dial server:", *address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client := pb.NewCPUServiceClient(conn)
	stream, err := client.Bidirectional(ctx)
	if err != nil {
		log.Panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	recvHandler := func() {
		for {
			if rsp, err := stream.Recv(); err == nil {
				log.Printf("Received Response: %v", core.ToJSONString(rsp))
			} else if err == io.EOF {
				log.Println("All Responses Have Received.")
				waitGroup.Done()
				return
			} else {
				log.Println("Error:", err, core.CurFuncDesc())
				waitGroup.Done()
				return
			}
		}
	}

	waitGroup.Add(1)
	sendHandler := func() {
		for i := 1; i <= 20; i++ {
			req := &pb.BidirectionalRequest{
				Id:    uuid.NewString(),
				Value: rand.Int63(),
			}

			log.Printf("(%v)Sending Request: %v", i, core.ToJSONString(req))
			if err := stream.Send(req); err != nil {
				waitGroup.Done()
				log.Panic(err)
				return
			}

			// Because the network speed on local host is very fast,
			// let this go rountine sleep a while after sending each request
			// to test/show the async result.
			time.Sleep(time.Microsecond * 100)
		}

		if err := stream.CloseSend(); err != nil {
			waitGroup.Done()
			log.Panic(err)
			return
		}

		log.Println("All Requests Have Been Sent.")
		waitGroup.Done()
	}

	go recvHandler()
	go sendHandler()
	waitGroup.Wait()
}
