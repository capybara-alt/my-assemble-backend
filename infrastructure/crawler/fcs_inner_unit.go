package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type fcsInnerUnit struct{}

func NewFcsInnerUnit() repository.ExternalInnerUnit {
	return &fcsInnerUnit{}
}

func (*fcsInnerUnit) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.FCS_PAGE, model.FCS_INNER_UNIT_TYPE)
}
