package handler

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

type InnerUnitHandler struct {
	repo    repository.InnerUnit
	usecase server.InnserUnitUsecase
}

func NewInnerUnitHandler(repo repository.InnerUnit) *InnerUnitHandler {
	return &InnerUnitHandler{
		repo:    repo,
		usecase: *server.NewInnserUnitUsecase(repo),
	}
}

func (i *InnerUnitHandler) GetInnerUnit(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetInnerUnitRequest],
) (*connect.Response[myassemblyv1.GetInnerUnitResponse], error) {
	resp, err := i.usecase.GetInnerUnit(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetInnerUnitResponse](resp), nil
}

func (i *InnerUnitHandler) GetInnerUnitList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[myassemblyv1.GetInnerUnitListResponse], error) {
	resp, err := i.usecase.GetInnerUnitList(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetInnerUnitListResponse](resp), nil
}
