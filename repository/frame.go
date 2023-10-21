package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Frame interface {
	UpsertBatch(context.Context, []model.Frame) error
	Get(context.Context, string) (*model.Frame, error)
	GetAll(context.Context) ([]model.Frame, error)
}

type ExternalFrame interface {
	Fetch() (model.CrawlResultJSON, error)
}
