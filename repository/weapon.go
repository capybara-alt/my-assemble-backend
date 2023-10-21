package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Weapon interface {
	UpsertBatch(context.Context, []model.Weapon) error
	Get(context.Context, string) (*model.Weapon, error)
	GetAll(context.Context) ([]model.Weapon, error)
}

type ExternalWeapon interface {
	Fetch() (model.CrawlResultJSON, error)
}
