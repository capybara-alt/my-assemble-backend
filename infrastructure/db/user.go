package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type user struct{}

func NewUser() repository.User {
	return &user{}
}

func (u *user) Create(ctx context.Context, user *model.User) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil

}

func (u *user) Get(ctx context.Context, key string) (*model.User, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	user := &model.User{}
	err := db.Where("id = ?", key).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}
