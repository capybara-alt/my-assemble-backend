package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type ExpansionCrawler struct {
	ICrawler
}

func NewExpansionCrawler() ICrawler {
	return &ExpansionCrawler{
		ICrawler: NewAbstractCrawler(config.EXPANSION_PAGE, model.EXPANSION_TYPE),
	}
}

func (c *ExpansionCrawler) JobName() string {
	return "ExpansionWeaponCrawlerJob"
}
