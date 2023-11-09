package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"playground/grpc/core"
	"playground/grpc/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type CPUService struct {
	ServiceBase
	pb.UnimplementedCPUServiceServer
}

func NewCPUService() *CPUService {
	return &CPUService{}
}

func (s *CPUService) CreateProcessor(
	ctx context.Context,
	in *pb.CreateProcessorRequest,
) (
	*pb.CreateProcessorResponse,
	error,
) {
	// If you want to test time out
	// time.Sleep(3 * time.Second)
	// if err := ctx.Err(); err != nil {
	// 	log.Printf("Error: %v, %v", err, core.CurFuncDesc())
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	log.Println(core.CurFuncDesc())

	if in.Cpu == nil {
		return nil, status.Error(codes.InvalidArgument, "Request.Cpu is nil")
	}

	return &pb.CreateProcessorResponse{
		Request: in,
		Cpuid:   uuid.NewString(),
	}, nil
}

func (s *CPUService) Query(in *pb.QueryRequest, stream pb.CPUService_QueryServer) error {
	log.Println("Request:", core.ToJSONString(in), core.CurFuncDesc())
	list := []*pb.CPU{
		{BrandName: "Intel"},
		{BrandName: "AMD"},
		{BrandName: "Apple"},
		{BrandName: "NVIDIA"},
	}

	for _, v := range list {
		if err := stream.Send(&pb.QueryResponse{
			Cpu: proto.Clone(v).(*pb.CPU),
		}); err != nil {
			return s.logError(fmt.Errorf("Error:%v, %v", err, core.CurFuncDesc()))
		} else {
			log.Println("Send CPU:", core.ToJSONString(v), core.CurFuncDesc())
		}
	}

	return nil
}

func (s *CPUService) UploadImage(stream pb.CPUService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return s.logError(status.Errorf(codes.Unknown, "Error: %v, %v", err, core.CurFuncDesc()))
	}

	log.Println("Request:", core.ToJSONString(req), core.CurFuncDesc())
	info := req.Data.(*pb.UploadImageRequest_Info).Info

	imageData := bytes.Buffer{}
	log.Println("Start receiving image data.", core.CurFuncDesc())
	for {
		if err := stream.Context().Err(); err != nil {
			return s.logError(status.Errorf(codes.Internal, "Error:%v, %v", err, core.CurFuncDesc()))
		}

		if req, err := stream.Recv(); err == io.EOF {
			log.Println("Finish receiving image data:", imageData.Len(), core.ToJSONString(info))
			break
		} else if err != nil {
			return s.logError(status.Errorf(codes.Unknown, "Error: %v, %v", err, core.CurFuncDesc()))
		} else {
			imageData.Write(req.GetChunk())
			log.Println("Get Chunk Size:", len(req.GetChunk()))
		}

		// If you want to test timeout.
		// time.Sleep(time.Second)
	}

	if err := stream.SendAndClose(&pb.UploadImageResponse{
		Info:      info,
		ImageId:   uuid.NewString(),
		ImageSize: uint32(imageData.Len()),
	}); err != nil {
		return s.logError(status.Errorf(codes.Unknown, "Error: %v, %v", err, core.CurFuncDesc()))
	}

	return nil
}

func (s *CPUService) Bidirectional(stream pb.CPUService_BidirectionalServer) error {
	var counter uint32
	for {
		if err := stream.Context().Err(); err != nil {
			return s.logError(status.Errorf(codes.Internal, "Error:%v, %v", err, core.CurFuncDesc()))
		}

		if req, err := stream.Recv(); err == io.EOF {
			log.Println("Finish receiving streaming data")
			break
		} else if err != nil {
			return s.logError(status.Errorf(codes.Unknown, "Error: %v, %v", err, core.CurFuncDesc()))
		} else {
			counter++
			log.Printf("(%v)Get Request: %v", counter, core.ToJSONString(req))

			rsp := &pb.BidirectionalResponse{
				Req:    req,
				Number: counter,
			}

			log.Printf("(%v)Send Response: %v", counter, core.ToJSONString(rsp))
			if err := stream.Send(rsp); err != nil {
				return s.logError(status.Errorf(codes.Unknown, "Error: %v, %v", err, core.CurFuncDesc()))
			}
		}
	}

	return nil
}

func (s *CPUService) DemoPermission(ctx context.Context, req *pb.DemoPermissionRequest) (*pb.DemoPermissionResponse, error) {
	return &pb.DemoPermissionResponse{
		Request: req,
		Status:  http.StatusOK,
	}, nil
}

func (s *CPUService) DemoTokenPermission(ctx context.Context, req *pb.DemoPermissionRequest) (*pb.DemoPermissionResponse, error) {
	return &pb.DemoPermissionResponse{
		Request: req,
		Status:  http.StatusOK,
	}, nil
}
