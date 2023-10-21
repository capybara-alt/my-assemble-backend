package server

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/repository"
)

type FrameUsecase struct {
	repo repository.Frame
}

func NewFrameUsecase(repo repository.Frame) *FrameUsecase {
	return &FrameUsecase{
		repo: repo,
	}
}

func (e *FrameUsecase) GetFrame(ctx context.Context, req *myassemblyv1.GetFrameRequest) (*myassemblyv1.GetFrameResponse, error) {
	frame, err := e.repo.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetFrameResponse{Item: frame.ToPB()}, nil
}

func (e *FrameUsecase) GetFrameList(ctx context.Context) (*myassemblyv1.GetFrameListResponse, error) {
	frameList, err := e.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*myassemblyv1.Frame, len(frameList))
	for i, frame := range frameList {
		resp[i] = frame.ToPB()
	}

	return &myassemblyv1.GetFrameListResponse{Items: resp}, nil
}
