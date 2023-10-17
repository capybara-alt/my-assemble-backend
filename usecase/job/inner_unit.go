package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type InnerUnitJob struct {
	dbRepo       repository.InnerUnit
	externalRepo []repository.ExternalInnerUnit
	convertor    convert.IConvertor[model.InnerUnit]
	logger       *slog.Logger
}

func NewInnerUnitJob(
	dbRepo repository.InnerUnit,
	externalRepo []repository.ExternalInnerUnit,
	convertor convert.IConvertor[model.InnerUnit],
	logger *slog.Logger) ICrawlJobUsecase {
	return &InnerUnitJob{
		dbRepo:       dbRepo,
		externalRepo: externalRepo,
		convertor:    convertor,
		logger:       logger,
	}
}

func (c *InnerUnitJob) Execute(ctx context.Context) {
	models := []model.InnerUnit{}

	for _, repo := range c.externalRepo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "detail", err)
		}
		c.logger.Info("Crawl successful")
		innerUnitList, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, innerUnitList...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.dbRepo.UpsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
