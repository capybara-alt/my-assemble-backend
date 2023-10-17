package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type ValidationUnitSchema interface {
	GetValidationSchema(context.Context, model.PrimaryUnitType) (model.ValidationUnitSchemaList, error)
}
