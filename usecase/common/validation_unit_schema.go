package common

import (
	"context"
	"log/slog"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type ValidationSchema struct {
	repo            repository.ValidationUnitSchema
	logger          *slog.Logger
	weaponSchema    model.ValidationUnitSchemaList
	frameSchema     model.ValidationUnitSchemaList
	innerUnitSchema model.ValidationUnitSchemaList
	expansionSchema model.ValidationUnitSchemaList
}

func NewValidationSchema(repo repository.ValidationUnitSchema, logger *slog.Logger) *ValidationSchema {
	return &ValidationSchema{
		repo:   repo,
		logger: logger,
	}
}

// TODO: Delete this function
func (v *ValidationSchema) Execute(ctx context.Context) {
	weaponSchema, err := v.repo.GetValidationSchema(ctx, model.WEAPON)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.weaponSchema = weaponSchema

	frameSchema, err := v.repo.GetValidationSchema(ctx, model.FRAME)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.frameSchema = frameSchema

	innerUnitSchema, err := v.repo.GetValidationSchema(ctx, model.INNER_UNITS)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.innerUnitSchema = innerUnitSchema

	expansionSchema, err := v.repo.GetValidationSchema(ctx, model.EXPANSION)
	if err != nil {
		v.logger.Error("Failed to fetch validation schema", "detail", err)
	}
	v.logger.Info("Fetch validation schema succeed")
	v.expansionSchema = expansionSchema
}

func (v *ValidationSchema) GetValidationUnitSchema(
	ctx context.Context,
	req *myassemblyv1.GetValidationUnitSchemaRequest,
) (*myassemblyv1.GetValidationUnitSchemaResponse, error) {
	vus, err := v.repo.GetValidationSchema(ctx, model.PrimaryUnitType(req.UnitType))
	if err != nil {
		return nil, err
	}

	return &myassemblyv1.GetValidationUnitSchemaResponse{Items: vus.ToPB()}, nil
}

func (v *ValidationSchema) GetWeaponSchema() model.ValidationUnitSchemaList {
	return v.weaponSchema
}

func (v *ValidationSchema) GetFrameSchema() model.ValidationUnitSchemaList {
	return v.frameSchema
}

func (v *ValidationSchema) GetInnerUnitsSchema() model.ValidationUnitSchemaList {
	return v.innerUnitSchema
}

func (v *ValidationSchema) GetExpansionSchema() model.ValidationUnitSchemaList {
	return v.expansionSchema
}
