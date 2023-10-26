package handler

import (
	"context"
	"log/slog"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	repo    repository.User
	usecase server.UserUsecase
	logger  *slog.Logger
}

func NewUserHandler(repo repository.User, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		repo:    repo,
		usecase: *server.NewUserUsecase(repo),
		logger:  logger,
	}
}

func (u *UserHandler) GetUser(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetUserRequest],
) (*connect.Response[myassemblyv1.GetUserResponse], error) {
	resp, err := u.usecase.GetUser(ctx, req.Msg)
	if err != nil {
		u.logger.Error("Get User Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetUserResponse](resp), nil
}

func (u *UserHandler) CreateUser(
	ctx context.Context,
	req *connect.Request[myassemblyv1.CreateUserRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := u.usecase.CreateUser(ctx, req.Msg)
	if err != nil {
		u.logger.Error("Create User Error", "detail", err)
		return nil, err
	}

	return &connect.Response[emptypb.Empty]{}, nil
}
