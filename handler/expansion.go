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

type ExpansionHandler struct {
	repo    repository.Expansion
	usecase server.ExpansionUsecase
	logger  *slog.Logger
}

func NewExpansionHandler(repo repository.Expansion, logger *slog.Logger) *ExpansionHandler {
	return &ExpansionHandler{
		repo:    repo,
		usecase: *server.NewExpansionUsecase(repo),
		logger:  logger,
	}
}

func (e *ExpansionHandler) GetExpansion(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetExpansionRequest],
) (*connect.Response[myassemblyv1.GetExpansionResponse], error) {
	resp, err := e.usecase.GetExpansion(ctx, req.Msg)
	if err != nil {
		e.logger.Error("Get Expansion Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetExpansionResponse](resp), nil
}

func (e *ExpansionHandler) GetExpansionList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[myassemblyv1.GetExpansionListResponse], error) {
	resp, err := e.usecase.GetExpansionList(ctx)
	if err != nil {
		e.logger.Error("Get Expansion Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetExpansionListResponse](resp), nil
}
