package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type extraWeapon struct{}

func NewExtraWeapon() repository.ExternalWeapon {
	return &extraWeapon{}
}

func (*extraWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.EXTRA_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE)
}
