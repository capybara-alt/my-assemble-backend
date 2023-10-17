package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Weapon interface {
	InsertBatch(context.Context, []model.Weapon) error
}

type ExternalWeapon interface {
	Fetch() (model.CrawlResultJSON, error)
}
