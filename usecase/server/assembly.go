package server

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type AssemblyUsecase struct {
	repo repository.Assembly
}

func NewAssemblyUsecase(repo repository.Assembly) *AssemblyUsecase {
	return &AssemblyUsecase{
		repo: repo,
	}
}

func (a *AssemblyUsecase) GetAssembly(ctx context.Context, req *myassemblyv1.GetAssemblyRequest) (*myassemblyv1.GetAssemblyResponse, error) {
	assembly, err := a.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetAssemblyResponse{Item: assembly.ToPB()}, nil
}

func (a *AssemblyUsecase) GetAssemblyList(ctx context.Context, req *myassemblyv1.GetAssemblyListRequest) (*myassemblyv1.GetAssemblyListResponse, error) {
	assemblyList, err := a.repo.GetList(ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	resp := make([]*myassemblyv1.Assembly, len(assemblyList))
	for i, assembly := range assemblyList {
		resp[i] = assembly.ToPB()
	}

	return &myassemblyv1.GetAssemblyListResponse{Items: resp}, nil
}

func (a *AssemblyUsecase) CreateAssembly(ctx context.Context, req *myassemblyv1.CreateAssemblyRequest) error {
	assembly := &model.Assembly{}
	err := a.repo.Create(ctx, *assembly.FromPB(req.Item))
	if err != nil {
		return err
	}

	return nil
}

func (a *AssemblyUsecase) UpdateAssembly(ctx context.Context, req *myassemblyv1.UpdateAssemblyRequest) (*myassemblyv1.UpdateAssemblyResponse, error) {
	assembly := &model.Assembly{}
	item, err := a.repo.Update(ctx, *assembly.FromPB(req.Item))
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.UpdateAssemblyResponse{Item: item.ToPB()}, nil
}
