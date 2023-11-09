package main

import (
	"context"
	"flag"
	"log"
	"playground/grpc/core"
	"playground/grpc/pb"

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

	// If you don't want to test time out
	ctx := context.Background()

	// If you want to test time out
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()

	client := pb.NewCPUServiceClient(conn)
	res, err := client.CreateProcessor(
		ctx,
		&pb.CreateProcessorRequest{
			Cpu: &pb.CPU{},
		},
	)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Response:", core.ToJSONString(res))
}
