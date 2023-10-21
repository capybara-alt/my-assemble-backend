package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type InnerUnit interface {
	UpsertBatch(context.Context, []model.InnerUnit) error
	Get(context.Context, string) (*model.InnerUnit, error)
	GetAll(context.Context) ([]model.InnerUnit, error)
}

type ExternalInnerUnit interface {
	Fetch() (model.CrawlResultJSON, error)
}
