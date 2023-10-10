package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type ExpansionJob struct {
	crawlers  []crawler.ICrawler
	repo      repository.Expansion
	convertor convert.IConvertor[model.Expansion]
	logger    *slog.Logger
}

func NewExpansionJob(
	crawlers []crawler.ICrawler,
	repo repository.Expansion,
	convertor convert.IConvertor[model.Expansion],
	logger *slog.Logger) ICrawlJobUsecase {
	return &ExpansionJob{
		crawlers:  crawlers,
		repo:      repo,
		convertor: convertor,
		logger:    logger,
	}
}

func (c *ExpansionJob) Execute(ctx context.Context) {
	models := []model.Expansion{}

	for _, crawler := range c.crawlers {
		results, err := crawler.Crawl()
		if err != nil {
			c.logger.Error("Crawl failed", "JobName", crawler.JobName(), "detail", err)
		}
		c.logger.Info("Crawl successful", "JobName", crawler.JobName())
		expansion_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "JobName", crawler.JobName(), "detail", err)
		} else {
			models = append(models, expansion_list...)
			c.logger.Info("Convert successful", "JobName", crawler.JobName())
		}
	}

	if err := c.repo.InsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
