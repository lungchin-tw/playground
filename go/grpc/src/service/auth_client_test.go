package service_test

import (
	"context"
	"net/http"
	"playground/grpc/core"
	"playground/grpc/pb"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	t.Parallel()

	address := startupTestServer(t)
	signup, login := signupThenLogin(t, address)
	user_id := signup.GetReq().UserId
	token := login.Token

	client := newTestCPUServiceClientWithUerToken(t, address, user_id, token)
	rsp, err := client.DemoTokenPermission(
		context.Background(),
		&pb.DemoPermissionRequest{
			UserId: user_id,
			Token:  token,
		},
	)

	t.Log("DemoTokenPermission Response:", core.ToJSONString(rsp))
	assert.NoError(t, err)
	assert.True(t, (http.StatusOK == rsp.Status))
}

func signupThenLogin(t *testing.T, address string) (*pb.NewUserSignupResponse, *pb.UserLoginResponse) {
	client := newTestAuthServiceClient(t, address)
	rspSignup, err := client.Signup(
		context.Background(),
		&pb.NewUserSignupRequest{
			UserId:   "User-" + uuid.NewString(),
			Password: "PW-" + uuid.NewString(),
		},
	)

	t.Log("Signup Response:", core.ToJSONString(rspSignup))
	assert.NoError(t, err)
	assert.True(t, (http.StatusOK == rspSignup.Status))

	rspLogin, err := client.Login(
		context.Background(),
		&pb.UserLoginRequest{
			UserId:   rspSignup.GetReq().GetUserId(),
			Password: rspSignup.GetReq().GetPassword(),
		},
	)

	t.Log("Login Response:", core.ToJSONString(rspLogin))
	assert.NoError(t, err)
	assert.True(t, (http.StatusOK == rspLogin.Status))
	return rspSignup, rspLogin
}
