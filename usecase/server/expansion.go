package server

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/repository"
)

type ExpansionUsecase struct {
	repo repository.Expansion
}

func NewExpansionUsecase(repo repository.Expansion) *ExpansionUsecase {
	return &ExpansionUsecase{
		repo: repo,
	}
}

func (e *ExpansionUsecase) GetExpansion(ctx context.Context, req *myassemblyv1.GetExpansionRequest) (*myassemblyv1.GetExpansionResponse, error) {
	expansion, err := e.repo.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetExpansionResponse{
		Item: expansion.ToPB(),
	}, nil
}

func (e *ExpansionUsecase) GetExpansionList(ctx context.Context) (*myassemblyv1.GetExpansionListResponse, error) {
	expansionList, err := e.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	respList := make([]*myassemblyv1.Expansion, len(expansionList))
	for i, expansion := range expansionList {
		respList[i] = expansion.ToPB()
	}

	return &myassemblyv1.GetExpansionListResponse{
		Items: respList,
	}, nil
}
