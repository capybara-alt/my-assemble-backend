package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type ExplodeWeaponCrawler struct {
	ICrawler
}

func NewExplodeWeaponCrawler() ICrawler {
	return &ExplodeWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.EXPLODE_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE),
	}
}

func (c *ExplodeWeaponCrawler) JobName() string {
	return "ExplodeWeaponCrawlJob"
}
