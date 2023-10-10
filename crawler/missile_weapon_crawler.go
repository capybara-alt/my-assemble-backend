package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type MissileWeaponCrawler struct {
	ICrawler
}

func NewMissileWeaponCrawler() ICrawler {
	return &MissileWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.MISSILE_WEAPON_PAGE, model.DEFAULT_WEAPON_BACK_UNIT_TYPE),
	}
}

func (c *MissileWeaponCrawler) JobName() string {
	return "MissileWeaponCrawlerJob"
}
