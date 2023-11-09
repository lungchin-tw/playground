package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"playground/grpc/core"
	"playground/grpc/pb"
	"runtime"
	"time"

	"github.com/google/uuid"
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
	stream, err := client.UploadImage(ctx)
	if err != nil {
		log.Panic(err)
	}

	_, path, _, ok := runtime.Caller(0)
	if ok == false {
		log.Panic(fmt.Errorf("runtime.Caller(0), Path=%v, OK=%v", path, ok))
	}

	imagePath := filepath.Join(filepath.Dir(path), "sample.png")
	log.Println("Image Path:", imagePath)

	f, err := os.Open(imagePath)
	defer f.Close()

	if err != nil {
		log.Panic(err)
	}

	if err := stream.Send(&pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				Device:    uuid.NewString(),
				ImageType: filepath.Ext(imagePath),
			},
		},
	}); err != nil {
		log.Panic(err)
	}

	reader := bufio.NewReader(f)
	buffer := make([]byte, 100<<10) // 1K

	for {
		if n, err := reader.Read(buffer); err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		} else {
			if err := stream.Send(&pb.UploadImageRequest{
				Data: &pb.UploadImageRequest_Chunk{
					Chunk: buffer[:n],
				},
			}); err != nil {
				log.Panic(err)
			}
		}
	}

	if rsp, err := stream.CloseAndRecv(); err != nil {
		log.Panic(err)
	} else {
		log.Println("Response:", core.ToJSONString(rsp))
	}
}
