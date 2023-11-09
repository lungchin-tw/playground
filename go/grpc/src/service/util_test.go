package service_test

import (
	"context"
	"net"
	"playground/grpc/core"
	"playground/grpc/pb"
	"playground/grpc/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func startupTestServer(t *testing.T) string {
	authService := service.NewAuthService()
	assert.NotNil(t, authService)

	interceptor := service.NewServerAuthInterceptor(
		map[string]bool{
			"/CPUService/DemoPermission":      false,
			"/CPUService/DemoTokenPermission": true,
		},
		authService.Secret(),
	)

	cpuService := service.NewCPUService()
	assert.NotNil(t, cpuService)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	assert.NotNil(t, grpcServer)

	pb.RegisterCPUServiceServer(grpcServer, cpuService)
	pb.RegisterAuthServiceServer(grpcServer, authService)

	listener, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)
	assert.NotNil(t, listener)
	assert.NotEmpty(t, listener.Addr())
	t.Log("Address:", listener.Addr().String())

	go grpcServer.Serve(listener)
	return listener.Addr().String()
}

func newTestAuthServiceClient(t *testing.T, serverAddress string) pb.AuthServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	assert.NoError(t, err)
	return pb.NewAuthServiceClient(conn)
}

func newTestCPUServiceClient(t *testing.T, serverAddress string) pb.CPUServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	assert.NoError(t, err)
	return pb.NewCPUServiceClient(conn)
}

func newTestCPUServiceClientWithUerToken(t *testing.T, serverAddress string, user_id string, token string) pb.CPUServiceClient {
	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			func(
				ctx context.Context,
				method string,
				req,
				reply interface{},
				cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker,
				opts ...grpc.CallOption,
			) error {
				t.Log(core.CurFuncDesc())
				t.Logf("Method:%v, Req:%v, Reply:%v", method, core.ToJSONString(req), core.ToJSONString(reply))

				ctx = metadata.AppendToOutgoingContext(ctx, service.FN_USER_ID, user_id)
				ctx = metadata.AppendToOutgoingContext(ctx, service.FN_AUTH_TOKEN, token)
				return invoker(ctx, method, req, reply, cc, opts...)
			},
		),
	}

	cc, err := grpc.Dial(serverAddress, options...)
	assert.NoError(t, err)
	return pb.NewCPUServiceClient(cc)
}
