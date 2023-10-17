package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type ExpansionJob struct {
	db_repo       repository.Expansion
	external_repo []repository.ExternalExpansion
	convertor     convert.IConvertor[model.Expansion]
	logger        *slog.Logger
}

func NewExpansionJob(
	db_repo repository.Expansion,
	external_repo []repository.ExternalExpansion,
	convertor convert.IConvertor[model.Expansion],
	logger *slog.Logger) ICrawlJobUsecase {
	return &ExpansionJob{
		db_repo:       db_repo,
		external_repo: external_repo,
		convertor:     convertor,
		logger:        logger,
	}
}

func (c *ExpansionJob) Execute(ctx context.Context) {
	models := []model.Expansion{}

	for _, repo := range c.external_repo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "JobName", "detail", err)
		}
		c.logger.Info("Crawl successful")
		expansion_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, expansion_list...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.db_repo.InsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
