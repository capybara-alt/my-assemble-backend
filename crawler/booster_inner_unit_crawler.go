package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type BoosterInnerUnitCrawler struct {
	ICrawler
}

func NewBoosterInnerUnitCrawler() ICrawler {
	return &BoosterInnerUnitCrawler{
		ICrawler: NewAbstractCrawler(config.BOOSTER_PAGE, model.BOOSTER_INNER_UNIT_TYPE),
	}
}

func (c *BoosterInnerUnitCrawler) JobName() string {
	return "BoosterInnerUnitCrawlerJob"
}
