package handler

import (
	"context"
	"log/slog"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/common"
)

type ValidationUnitSchemaHandler struct {
	repo    repository.ValidationUnitSchema
	usecase common.ValidationSchema
	logger  *slog.Logger
}

func NewValidationUnitSchemaHandler(
	repo repository.ValidationUnitSchema,
	logger *slog.Logger,
) *ValidationUnitSchemaHandler {
	return &ValidationUnitSchemaHandler{
		repo:    repo,
		usecase: *common.NewValidationSchema(repo, logger),
	}
}

func (v *ValidationUnitSchemaHandler) GetValidationUnitSchema(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetValidationUnitSchemaRequest],
) (*connect.Response[myassemblyv1.GetValidationUnitSchemaResponse], error) {
	resp, err := v.usecase.GetValidationUnitSchema(ctx, req.Msg)
	if err != nil {
		v.logger.Error("Get ValidationUnitSchema Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetValidationUnitSchemaResponse](resp), nil
}
