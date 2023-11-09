package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"playground/grpc/core"
	"playground/grpc/pb"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	FN_USER_ID     = "user_id"
	FN_AUTH_TOKEN  = "auth_token"
	AUTH_TOKEN_TTL = time.Second * 2
)

type User struct {
	UserID         string
	HashedPassword string
	Token          string
}

type AuthService struct {
	ServiceBase
	pb.UnimplementedAuthServiceServer
	UserRepo map[string]*User
	secret   string
	tokenTTL time.Duration
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserRepo: map[string]*User{},
		secret:   uuid.NewString(),
		tokenTTL: AUTH_TOKEN_TTL,
	}
}

func (s *AuthService) Secret() string {
	return s.secret
}

func (s *AuthService) TokenTTL() time.Duration {
	return s.tokenTTL
}

func (s *AuthService) Signup(
	ctx context.Context,
	req *pb.NewUserSignupRequest,
) (
	*pb.NewUserSignupResponse,
	error,
) {
	hashpw, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return &pb.NewUserSignupResponse{
				Req:     req,
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			},
			s.logError(status.Errorf(codes.Internal, "Error: %v, %v", err, core.CurFuncDesc()))
	}

	s.UserRepo[req.GetUserId()] = &User{
		UserID:         req.GetUserId(),
		HashedPassword: string(hashpw),
	}

	return &pb.NewUserSignupResponse{
		Req:    req,
		Status: http.StatusOK,
	}, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	req *pb.UserLoginRequest,
) (
	*pb.UserLoginResponse,
	error,
) {
	log.Println("Request:", core.ToJSONString(req))
	if v, ok := s.UserRepo[req.GetUserId()]; ok == false {
		err := s.logError(status.Errorf(codes.NotFound, "%v, %v", req.GetUserId(), core.CurFuncDesc()))
		return &pb.UserLoginResponse{
			Status:  http.StatusUnauthorized,
			Message: fmt.Sprintf("User %v Not Found", req.GetUserId()),
		}, err
	} else if err := bcrypt.CompareHashAndPassword([]byte(v.HashedPassword), []byte(req.GetPassword())); err != nil {
		err := s.logError(status.Errorf(codes.InvalidArgument, "Error:%v, %v", err, core.CurFuncDesc()))
		return &pb.UserLoginResponse{
			Status:  http.StatusUnauthorized,
			Message: "Incorrect Password",
		}, err
	}

	token, err := core.GenerateToken(req.GetUserId(), s.Secret(), s.TokenTTL())
	if err != nil {
		err := s.logError(status.Errorf(codes.Internal, "Error:%v, %v", err, core.CurFuncDesc()))
		return &pb.UserLoginResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, err
	}

	log.Println("Token:", token)

	return &pb.UserLoginResponse{
		Token:  token,
		Status: http.StatusOK,
	}, nil
}

type ServerAuthInterceptor struct {
	ServiceBase
	accessPermission map[string]bool
	secret           string
}

func NewServerAuthInterceptor(permission map[string]bool, secret string) *ServerAuthInterceptor {
	return &ServerAuthInterceptor{
		accessPermission: permission,
		secret:           secret,
	}
}

func (i *ServerAuthInterceptor) Secret() string {
	return i.secret
}

func (i *ServerAuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println(core.CurFuncDesc())
		log.Println(core.ToJSONString(info))

		if err := i.authorize(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i *ServerAuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println(core.CurFuncDesc())
		log.Println(core.ToJSONString(info))
		return handler(srv, ss)
	}
}

func (i *ServerAuthInterceptor) authorize(ctx context.Context, method string) error {
	// To demo the authentication more easily,
	// if the method is not set in the accessPermission, allow its authentication.
	if v, ok := i.accessPermission[method]; ok == false {
		return nil
	} else if v == false {
		return i.logError(status.Errorf(codes.PermissionDenied, "%v, %v, %v", method, v, core.CurFuncDesc()))
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return i.logError(status.Errorf(codes.Unauthenticated, "%v, %v", method, core.CurFuncDesc()))
	}

	log.Println("MetaData:", core.ToJSONString(md))

	token, ok := md[FN_AUTH_TOKEN]
	if (ok == false) || (len(token) == 0) {
		return i.logError(status.Errorf(codes.NotFound, "%v, %v", FN_AUTH_TOKEN, core.CurFuncDesc()))
	}

	log.Println("Token:", core.ToJSONString(md))
	claims, err := core.Verify(token[0], i.Secret())
	if err != nil {
		return i.logError(status.Errorf(codes.Unauthenticated, "%v, %v, %v", method, err, core.CurFuncDesc()))
	}

	user_id, ok := md[FN_USER_ID]
	if (ok == false) || (len(user_id) == 0) {
		return i.logError(status.Errorf(codes.NotFound, "%v, %v", FN_USER_ID, core.CurFuncDesc()))
	} else if claims.Id != user_id[0] {
		return i.logError(status.Errorf(codes.Unauthenticated, "%v, %v", user_id[0], core.CurFuncDesc()))
	}

	log.Println("Claims:", core.ToJSONString(claims))
	return nil
}
