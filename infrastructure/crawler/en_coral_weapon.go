package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type enCoralWeapon struct{}

func NewEnCoralWeapon() repository.ExternalWeapon {
	return &enCoralWeapon{}
}

func (*enCoralWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.EN_CORAL_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE)
}
