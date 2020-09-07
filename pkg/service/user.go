package service

import (
	context "context"
	"errors"

	"github.com/bmf-san/golang-grpc-example/pkg/proto/user"
)

type UserService struct{}

func (s *UserService) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	switch req.Type {
	case "admin":
		return &user.GetUserResponse{
			Name: "admin_user",
		}, nil
	case "general":
		return &user.GetUserResponse{
			Name: "general_user",
		}, nil
	}

	return nil, errors.New("No user")
}
