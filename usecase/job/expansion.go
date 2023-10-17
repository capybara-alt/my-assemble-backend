package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type ExpansionJob struct {
	dbRepo       repository.Expansion
	externalRepo []repository.ExternalExpansion
	convertor    convert.IConvertor[model.Expansion]
	logger       *slog.Logger
}

func NewExpansionJob(
	dbRepo repository.Expansion,
	externalRepo []repository.ExternalExpansion,
	convertor convert.IConvertor[model.Expansion],
	logger *slog.Logger) ICrawlJobUsecase {
	return &ExpansionJob{
		dbRepo:       dbRepo,
		externalRepo: externalRepo,
		convertor:    convertor,
		logger:       logger,
	}
}

func (c *ExpansionJob) Execute(ctx context.Context) {
	models := []model.Expansion{}

	for _, repo := range c.externalRepo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "JobName", "detail", err)
		}
		c.logger.Info("Crawl successful")
		expansionList, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, expansionList...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.dbRepo.UpsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
