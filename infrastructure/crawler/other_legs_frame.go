package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type otherLegsFrame struct{}

func NewOtherLegsFrame() repository.ExternalFrame {
	return &otherLegsFrame{}
}

func (*otherLegsFrame) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.LEGS_OTHER_PARTS_PAGE, model.LEGS_UNIT_TYPE)
}
