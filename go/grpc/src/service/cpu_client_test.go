package service_test

import (
	"bufio"
	"context"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"playground/grpc/core"
	"playground/grpc/pb"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnaryClientFail(t *testing.T) {
	t.Parallel()

	address := startupTestServer(t)
	client := newTestCPUServiceClient(t, address)
	assert.NotNil(t, client)

	res, err := client.CreateProcessor(
		context.Background(),
		&pb.CreateProcessorRequest{},
	)

	assert.Empty(t, res)
	assert.Error(t, err)
	t.Log("Error:", err)
}

func TestUnaryClientSucceed(t *testing.T) {
	t.Parallel()

	address := startupTestServer(t)
	client := newTestCPUServiceClient(t, address)
	assert.NotNil(t, client)

	res, err := client.CreateProcessor(
		context.Background(),
		&pb.CreateProcessorRequest{Cpu: &pb.CPU{}},
	)

	assert.NotNil(t, res)
	assert.NoError(t, err)
	t.Log("Response:", core.ToJSONString(res))
}

func TestRcvStream(t *testing.T) {
	t.Parallel()

	address := startupTestServer(t)
	client := newTestCPUServiceClient(t, address)
	assert.NotNil(t, client)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stream, err := client.Query(
		ctx,
		&pb.QueryRequest{Query: &pb.Query{Keyword: "*"}},
	)

	assert.NoError(t, err)

	for {
		if rsp, err := stream.Recv(); err == nil {
			t.Log("Received:", core.ToJSONString(rsp))
		} else if err == io.EOF {
			break
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestUploadStream(t *testing.T) {
	t.Parallel()

	address := startupTestServer(t)
	client := newTestCPUServiceClient(t, address)
	assert.NotNil(t, client)

	_, path, _, ok := runtime.Caller(0)
	assert.True(t, ok)

	imagePath := filepath.Join(filepath.Dir(path), "../cmd/client/uploadstream/sample.png")
	t.Log("Image Path:", imagePath)

	f, err := os.Open(imagePath)
	defer f.Close()
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stream, err := client.UploadImage(ctx)
	assert.NoError(t, err)

	assert.NoError(t, stream.Send(&pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				Device:    uuid.NewString(),
				ImageType: filepath.Ext(imagePath),
			},
		},
	}))

	reader := bufio.NewReader(f)
	buffer := make([]byte, 100<<10) // 1K

	for {
		if n, err := reader.Read(buffer); err == io.EOF {
			break
		} else {
			assert.NoError(t, err)
			assert.NoError(t, stream.Send(&pb.UploadImageRequest{
				Data: &pb.UploadImageRequest_Chunk{
					Chunk: buffer[:n],
				},
			}))
		}
	}

	rsp, err := stream.CloseAndRecv()
	assert.NoError(t, err)
	t.Log("Response:", core.ToJSONString(rsp))
}

func TestBidirectionalStream(t *testing.T) {
	t.Parallel()

	address := startupTestServer(t)
	client := newTestCPUServiceClient(t, address)
	assert.NotNil(t, client)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stream, err := client.Bidirectional(ctx)
	assert.NoError(t, err)

	rand.Seed(time.Now().UnixNano())

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	recvHandler := func(t *testing.T) {
		for {
			if rsp, err := stream.Recv(); err == nil {
				t.Logf("Received Response: %v", core.ToJSONString(rsp))
			} else if err == io.EOF {
				t.Log("All Responses Have Received.")
				waitGroup.Done()
				return
			} else {
				waitGroup.Done()
				assert.NoError(t, err)
				return
			}
		}
	}

	waitGroup.Add(1)
	sendHandler := func(t *testing.T) {
		for i := 1; i <= 20; i++ {
			req := &pb.BidirectionalRequest{
				Id:    uuid.NewString(),
				Value: rand.Int63(),
			}

			t.Logf("(%v)Sending Request: %v", i, core.ToJSONString(req))
			if err := stream.Send(req); err != nil {
				waitGroup.Done()
				assert.NoError(t, err)
				return
			}

			// Because the network speed on local host is very fast,
			// let this go rountine sleep a while after sending each request
			// to test/show the async result.
			time.Sleep(time.Microsecond * 100)
		}

		waitGroup.Done()
		assert.NoError(t, stream.CloseSend())
	}

	go recvHandler(t)
	go sendHandler(t)
	waitGroup.Wait()
}
