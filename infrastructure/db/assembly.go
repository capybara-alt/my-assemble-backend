package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type assembly struct{}

func NewAssembly() repository.Assembly {
	return &assembly{}
}

func (a *assembly) Create(ctx context.Context, assembly model.Assembly) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	if err := db.Save(&assembly).Error; err != nil {
		return err
	}
	return nil
}

func (a *assembly) Get(ctx context.Context, key int64) (*model.Assembly, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	assembly := &model.Assembly{}
	err := db.First(&assembly, key).Error
	if err != nil {
		return nil, err
	}

	return assembly, nil
}

func (a *assembly) GetList(ctx context.Context, key string) ([]model.Assembly, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	assemblies := []model.Assembly{}
	err := db.Where("user_uid = ?", key).Find(&assemblies).Error
	if err != nil {
		return nil, err
	}

	return assemblies, nil
}

func (a *assembly) Update(ctx context.Context, assembly model.Assembly) (*model.Assembly, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	if err := db.Save(&assembly).Error; err != nil {
		return nil, err
	}

	return &assembly, nil
}
