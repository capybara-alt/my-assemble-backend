package server

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type UserUsecase struct {
	repo repository.User
}

func NewUserUsecase(repo repository.User) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) GetUser(ctx context.Context, req *myassemblyv1.GetUserRequest) (*myassemblyv1.GetUserResponse, error) {
	user, err := u.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetUserResponse{Item: user.ToPB()}, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, req *myassemblyv1.CreateUserRequest) error {
	user := &model.User{}
	err := u.repo.Create(ctx, user.FromPB(req.Item))
	if err != nil {
		return err
	}

	return nil
}
