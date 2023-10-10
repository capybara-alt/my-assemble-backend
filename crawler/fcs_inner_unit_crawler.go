package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type FcsInnerUnitCrawler struct {
	ICrawler
}

func NewFcsInnerUnitCrawler() ICrawler {
	return &FcsInnerUnitCrawler{
		ICrawler: NewAbstractCrawler(config.FCS_PAGE, model.FCS_INNER_UNIT_TYPE),
	}
}

func (c *FcsInnerUnitCrawler) JobName() string {
	return "FcsInnerUnitCrawlerJob"
}
