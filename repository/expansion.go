package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Expansion interface {
	UpsertBatch(context.Context, []model.Expansion) error
	Get(context.Context, string) (*model.Expansion, error)
	GetAll(context.Context) ([]model.Expansion, error)
}

type ExternalExpansion interface {
	Fetch() (model.CrawlResultJSON, error)
}
