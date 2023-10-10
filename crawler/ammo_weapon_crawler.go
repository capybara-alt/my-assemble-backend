package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type AmmoWeaponCrawler struct {
	ICrawler
}

func NewAmmoWeaponCrawler() ICrawler {
	return &AmmoWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.AMMO_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE),
	}
}

func (c *AmmoWeaponCrawler) JobName() string {
	return "AmmoWeaponCrawlJob"
}
