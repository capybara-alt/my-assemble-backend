package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type explodeWeapon struct{}

func NewExplodeWeapon() repository.ExternalWeapon {
	return &explodeWeapon{}
}

func (*explodeWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.EXPLODE_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE)
}
