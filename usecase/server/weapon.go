package server

import (
	"context"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/repository"
)

type WeaponUsecase struct {
	repo repository.Weapon
}

func NewWeaponUsecase(repo repository.Weapon) *WeaponUsecase {
	return &WeaponUsecase{
		repo: repo,
	}
}

func (e *WeaponUsecase) GetWeapon(ctx context.Context, req *myassemblyv1.GetWeaponRequest) (*myassemblyv1.GetWeaponResponse, error) {
	weapon, err := e.repo.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetWeaponResponse{Item: weapon.ToPB()}, nil
}

func (e *WeaponUsecase) GetWeaponList(ctx context.Context) (*myassemblyv1.GetWeaponListResponse, error) {
	weaponList, err := e.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*myassemblyv1.Weapon, len(weaponList))
	for i, weapon := range weaponList {
		resp[i] = weapon.ToPB()
	}

	return &myassemblyv1.GetWeaponListResponse{Items: resp}, nil
}
