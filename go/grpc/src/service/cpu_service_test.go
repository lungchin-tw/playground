package service_test

import (
	"context"
	"playground/grpc/core"
	"playground/grpc/pb"
	"playground/grpc/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUnaryServerFail(t *testing.T) {
	t.Parallel()

	req := &pb.CreateProcessorRequest{}
	server := service.NewCPUService()
	res, err := server.CreateProcessor(context.Background(), req)
	assert.Error(t, err)
	st, ok := status.FromError(err)
	t.Log("Status:", core.ToJSONString(st.Proto()))
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.Nil(t, res)
}

func TestUnaryServerSucceed(t *testing.T) {
	t.Parallel()

	req := &pb.CreateProcessorRequest{Cpu: &pb.CPU{}}
	server := service.NewCPUService()
	res, err := server.CreateProcessor(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	t.Log("Response:", core.ToJSONString(res))
}
