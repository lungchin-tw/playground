package main

import (
	"context"
	"flag"
	"io"
	"log"
	"playground/grpc/core"
	"playground/grpc/pb"
	"time"

	"google.golang.org/grpc"
)

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
	stream, err := client.Query(
		ctx,
		&pb.QueryRequest{Query: &pb.Query{Keyword: "*"}},
	)

	if err != nil {
		log.Panic(err)
	}

	for {
		if rsp, err := stream.Recv(); err == nil {
			log.Println("Received:", core.ToJSONString(rsp))
		} else if err == io.EOF {
			break
		} else {
			log.Panic(err)
		}
	}
}
