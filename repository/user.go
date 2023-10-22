package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type User interface {
	Get(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) error
}
