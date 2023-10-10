package repository

import (
	"context"

	"github.com/capybara-alt/my-assemble/model"
)

type ValidationUnitSchema interface {
	GetValidationSchema(ctx context.Context, unit_type model.PrimaryUnitType) (model.ValidationUnitSchemaList, error)
}
