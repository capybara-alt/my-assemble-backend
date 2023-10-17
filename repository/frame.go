package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type Frame interface {
	InsertBatch(context.Context, []model.Frame) error
}

type ExternalFrame interface {
	Fetch() (model.CrawlResultJSON, error)
}
