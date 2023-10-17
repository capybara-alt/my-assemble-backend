package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type coreFrame struct{}

func NewCoreFrame() repository.ExternalInnerUnit {
	return &coreFrame{}
}

func (*coreFrame) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.CORE_PARTS_PAGE, model.CORE_UNIT_TYPE)
}
