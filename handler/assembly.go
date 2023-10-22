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

type AssemblyHandler struct {
	repo    repository.Assembly
	usecase server.AssemblyUsecase
	logger  *slog.Logger
}

func NewAssemblyHandler(repo repository.Assembly, logger *slog.Logger) *AssemblyHandler {
	return &AssemblyHandler{
		repo:    repo,
		usecase: *server.NewAssemblyUsecase(repo),
		logger:  logger,
	}
}

func (a *AssemblyHandler) GetAssembly(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetAssemblyRequest],
) (*connect.Response[myassemblyv1.GetAssemblyResponse], error) {
	resp, err := a.usecase.GetAssembly(ctx, req.Msg)
	if err != nil {
		a.logger.Error("Get Assembly Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetAssemblyResponse](resp), nil
}

func (a *AssemblyHandler) GetAssemblyList(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetAssemblyListRequest],
) (*connect.Response[myassemblyv1.GetAssemblyListResponse], error) {
	resp, err := a.usecase.GetAssemblyList(ctx, req.Msg)
	if err != nil {
		a.logger.Error("Get Assembly Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetAssemblyListResponse](resp), nil
}

func (a *AssemblyHandler) CreateAssembly(
	ctx context.Context,
	req *connect.Request[myassemblyv1.CreateAssemblyRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := a.usecase.CreateAssembly(ctx, req.Msg)
	if err != nil {
		a.logger.Error("Create Assembly Error", "detail", err)
		return nil, err
	}

	return &connect.Response[emptypb.Empty]{}, nil
}

func (a *AssemblyHandler) UpdateAssembly(
	ctx context.Context,
	req *connect.Request[myassemblyv1.UpdateAssemblyRequest],
) (*connect.Response[myassemblyv1.UpdateAssemblyResponse], error) {
	resp, err := a.usecase.UpdateAssembly(ctx, req.Msg)
	if err != nil {
		a.logger.Error("Update Assembly Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.UpdateAssemblyResponse](resp), nil
}
