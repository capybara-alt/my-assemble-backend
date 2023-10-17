package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type InnerUnitJob struct {
	db_repo       repository.InnerUnit
	external_repo []repository.ExternalInnerUnit
	convertor     convert.IConvertor[model.InnerUnit]
	logger        *slog.Logger
}

func NewInnerUnitJob(
	db_repo repository.InnerUnit,
	external_repo []repository.ExternalInnerUnit,
	convertor convert.IConvertor[model.InnerUnit],
	logger *slog.Logger) ICrawlJobUsecase {
	return &InnerUnitJob{
		db_repo:       db_repo,
		external_repo: external_repo,
		convertor:     convertor,
		logger:        logger,
	}
}

func (c *InnerUnitJob) Execute(ctx context.Context) {
	models := []model.InnerUnit{}

	for _, repo := range c.external_repo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "detail", err)
		}
		c.logger.Info("Crawl successful")
		inner_units_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, inner_units_list...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.db_repo.InsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
