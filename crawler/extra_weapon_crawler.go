package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type ExtraWeaponCrawler struct {
	ICrawler
}

func NewExtraWeaponCrawler() ICrawler {
	return &ExtraWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.EXTRA_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE),
	}
}

func (c *ExtraWeaponCrawler) JobName() string {
	return "ExtraWeaponCrawlerJob"
}
