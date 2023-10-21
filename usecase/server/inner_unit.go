package server

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/repository"
)

type InnserUnitUsecase struct {
	repo repository.InnerUnit
}

func NewInnserUnitUsecase(repo repository.InnerUnit) *InnserUnitUsecase {
	return &InnserUnitUsecase{
		repo: repo,
	}
}

func (e *InnserUnitUsecase) GetInnerUnit(ctx context.Context, req *myassemblyv1.GetInnerUnitRequest) (*myassemblyv1.GetInnerUnitResponse, error) {
	innerUnit, err := e.repo.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetInnerUnitResponse{Item: innerUnit.ToPB()}, nil
}

func (e *InnserUnitUsecase) GetInnerUnitList(ctx context.Context) (*myassemblyv1.GetInnerUnitListResponse, error) {
	innerUnitList, err := e.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*myassemblyv1.InnerUnit, len(innerUnitList))
	for i, innerUnit := range innerUnitList {
		resp[i] = innerUnit.ToPB()
	}

	return &myassemblyv1.GetInnerUnitListResponse{Items: resp}, nil
}
