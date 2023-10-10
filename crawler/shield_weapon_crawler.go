package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type ShieldWeaponCrawler struct {
	ICrawler
}

func NewShieldWeaponCrawler() ICrawler {
	return &ShieldWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.SHIELD_WEAPON_PAGE, model.SHIELD_WEAPON_UNIT_TYPE),
	}
}

func (c *ShieldWeaponCrawler) JobName() string {
	return "ShieldWeaponCrawlerJob"
}
