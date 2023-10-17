package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type InnerUnit interface {
	InsertBatch(context.Context, []model.InnerUnit) error
}

type ExternalInnerUnit interface {
	Fetch() (model.CrawlResultJSON, error)
}
