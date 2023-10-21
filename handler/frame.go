package handler

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FrameHandler struct {
	repo    repository.Frame
	usecase server.FrameUsecase
}

func NewFrameHandler(repo repository.Frame) *FrameHandler {
	return &FrameHandler{
		repo:    repo,
		usecase: *server.NewFrameUsecase(repo),
	}
}

func (f *FrameHandler) GetFrame(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetFrameRequest],
) (*connect.Response[myassemblyv1.GetFrameResponse], error) {
	resp, err := f.usecase.GetFrame(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetFrameResponse](resp), nil
}

func (f *FrameHandler) GetFrameList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[myassemblyv1.GetFrameListResponse], error) {
	resp, err := f.usecase.GetFrameList(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetFrameListResponse](resp), nil
}
