package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type InnerUnitJob struct {
	crawlers  []crawler.ICrawler
	repo      repository.InnerUnit
	convertor convert.IConvertor[model.InnerUnit]
	logger    *slog.Logger
}

func NewInnerUnitJob(
	crawlers []crawler.ICrawler,
	repo repository.InnerUnit,
	convertor convert.IConvertor[model.InnerUnit],
	logger *slog.Logger) ICrawlJobUsecase {
	return &InnerUnitJob{
		crawlers:  crawlers,
		repo:      repo,
		convertor: convertor,
		logger:    logger,
	}
}

func (c *InnerUnitJob) Execute(ctx context.Context) {
	models := []model.InnerUnit{}

	for _, crawler := range c.crawlers {
		results, err := crawler.Crawl()
		if err != nil {
			c.logger.Error("Crawl failed", "JobName", crawler.JobName(), "detail", err)
		}
		c.logger.Info("Crawl successful", "JobName", crawler.JobName())
		inner_units_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "JobName", crawler.JobName(), "detail", err)
		} else {
			models = append(models, inner_units_list...)
			c.logger.Info("Convert successful", "JobName", crawler.JobName())
		}
	}

	if err := c.repo.InsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
