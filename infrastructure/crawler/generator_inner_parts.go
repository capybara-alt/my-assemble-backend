package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type generatorInnerUnit struct{}

func NewGeneratorInnerUnit() repository.ExternalInnerUnit {
	return &generatorInnerUnit{}
}

func (*generatorInnerUnit) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.GENERATOR_PAGE, model.GENERATOR_INNER_UNIT_TYPE)
}
