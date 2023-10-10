package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type OrbitTaletDroneWeaponCrawler struct {
	ICrawler
}

func NewOrbitTaletDroneWeaponCrawler() ICrawler {
	return &OrbitTaletDroneWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.ORBIT_WEAPON_PAGE, model.DEFAULT_WEAPON_BACK_UNIT_TYPE),
	}
}

func (c *OrbitTaletDroneWeaponCrawler) JobName() string {
	return "OrbitTaletDroneWeaponCrawlerJob"
}
