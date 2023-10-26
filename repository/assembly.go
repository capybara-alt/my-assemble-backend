package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Assembly interface {
	Get(context.Context, int64) (*model.Assembly, error)
	GetList(context.Context, string) ([]model.Assembly, error)
	Create(context.Context, model.Assembly) error
	Update(context.Context, model.Assembly) (*model.Assembly, error)
}
