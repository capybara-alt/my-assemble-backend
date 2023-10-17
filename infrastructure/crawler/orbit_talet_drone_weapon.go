package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type orbitTaletDroneWeapon struct{}

func NewOrbitTaletDroneWeapon() repository.ExternalWeapon {
	return &orbitTaletDroneWeapon{}
}

func (*orbitTaletDroneWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.ORBIT_WEAPON_PAGE, model.DEFAULT_WEAPON_BACK_UNIT_TYPE)
}
