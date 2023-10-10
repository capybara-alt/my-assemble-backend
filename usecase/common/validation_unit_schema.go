package common

import (
	"context"
	"log/slog"

	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type ValidationSchema struct {
	repo               repository.ValidationUnitSchema
	logger             *slog.Logger
	weapon_schema      model.ValidationUnitSchemaList
	frame_schema       model.ValidationUnitSchemaList
	inner_units_schema model.ValidationUnitSchemaList
	expansion_schema   model.ValidationUnitSchemaList
}

func NewValidationSchema(repo repository.ValidationUnitSchema, logger *slog.Logger) *ValidationSchema {
	return &ValidationSchema{
		repo:   repo,
		logger: logger,
	}
}

func (v *ValidationSchema) Execute(ctx context.Context) {
	weapon_schema, err := v.repo.GetValidationSchema(ctx, model.WEAPON)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.weapon_schema = weapon_schema

	frame_schema, err := v.repo.GetValidationSchema(ctx, model.FRAME)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.frame_schema = frame_schema

	inner_units_schema, err := v.repo.GetValidationSchema(ctx, model.INNER_UNITS)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.inner_units_schema = inner_units_schema

	expansion_schema, err := v.repo.GetValidationSchema(ctx, model.EXPANSION)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.expansion_schema = expansion_schema
}

func (v *ValidationSchema) GetWeaponSchema() model.ValidationUnitSchemaList {
	return v.weapon_schema
}

func (v *ValidationSchema) GetFrameSchema() model.ValidationUnitSchemaList {
	return v.frame_schema
}

func (v *ValidationSchema) GetInnerUnitsSchema() model.ValidationUnitSchemaList {
	return v.inner_units_schema
}

func (v *ValidationSchema) GetExpansionSchema() model.ValidationUnitSchemaList {
	return v.expansion_schema
}
