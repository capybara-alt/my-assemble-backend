package convert

import "github.com/capybara-alt/my-assemble/model"

type IConvertor[T model.Frame | model.Weapon | model.InnerUnit | model.Expansion] interface {
	Convert(model.CrawlResultJSON) ([]T, error)
}
