package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"playground/grpc/core"
	"playground/grpc/pb"
	"playground/grpc/service"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	address := flag.String("address", "", "the server address")
	flag.Parse()
	log.Println("Dial server:", *address)

	cc, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewAuthServiceClient(cc)
	resSignup, err := client.Signup(
		context.Background(),
		&pb.NewUserSignupRequest{
			UserId:   "User-" + uuid.NewString(),
			Password: "PW-" + uuid.NewString(),
		},
	)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Signup Response:", core.ToJSONString(resSignup))

	if resSignup.Status != http.StatusOK {
		log.Panic(fmt.Errorf("Signup Response Status != http.StatusOK"))
	}

	resLogin, err := client.Login(
		context.Background(),
		&pb.UserLoginRequest{
			UserId:   resSignup.GetReq().GetUserId(),
			Password: resSignup.GetReq().GetPassword(),
		},
	)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Login Response:", core.ToJSONString(resLogin))

	if resLogin.Status != http.StatusOK {
		log.Panic(fmt.Errorf("Login Response Status != http.StatusOK"))
	}

	// If you want to test token expiration
	// time.Sleep(service.AUTH_TOKEN_TTL)

	callDemoTokenPermission(*address, resSignup.GetReq().GetUserId(), resLogin.GetToken())
}

func callDemoTokenPermission(address string, user_id string, token string) {
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
				log.Println(core.CurFuncDesc())
				log.Printf("Method:%v, Req:%v, Reply:%v", method, core.ToJSONString(req), core.ToJSONString(reply))

				ctx = metadata.AppendToOutgoingContext(ctx, service.FN_AUTH_TOKEN, token)
				ctx = metadata.AppendToOutgoingContext(ctx, service.FN_USER_ID, user_id)
				return invoker(ctx, method, req, reply, cc, opts...)
			},
		),
		grpc.WithStreamInterceptor(
			func(
				ctx context.Context,
				desc *grpc.StreamDesc,
				cc *grpc.ClientConn,
				method string,
				streamer grpc.Streamer,
				opts ...grpc.CallOption,
			) (
				grpc.ClientStream,
				error,
			) {
				log.Println(core.CurFuncDesc())
				log.Printf("Desc:%v, Method:%v", core.ToJSONString(desc), method)
				return streamer(ctx, desc, cc, method, opts...)
			},
		),
	}

	cc, err := grpc.Dial(address, options...)
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewCPUServiceClient(cc)
	rsp, err := client.DemoTokenPermission(
		context.Background(),
		&pb.DemoPermissionRequest{
			UserId: user_id,
			Token:  token,
		},
	)

	if err != nil {
		log.Panic(err)
	}

	log.Println("callDemoTokenPermission, Response:", core.ToJSONString(rsp))

	if rsp.GetStatus() != http.StatusOK {
		log.Panic(fmt.Errorf("callDemoTokenPermission Response Status != http.StatusOK"))
	}
}
