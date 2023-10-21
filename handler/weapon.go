package handler

import (
	"context"
	"log/slog"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WeaponHandler struct {
	repo    repository.Weapon
	usecase server.WeaponUsecase
	logger  *slog.Logger
}

func NewWeaponHandler(repo repository.Weapon, logger *slog.Logger) *WeaponHandler {
	return &WeaponHandler{
		repo:    repo,
		usecase: *server.NewWeaponUsecase(repo),
		logger:  logger,
	}
}

func (w *WeaponHandler) GetWeapon(
	ctx context.Context,
	req *connect.Request[myassemblyv1.GetWeaponRequest],
) (*connect.Response[myassemblyv1.GetWeaponResponse], error) {
	resp, err := w.usecase.GetWeapon(ctx, req.Msg)
	if err != nil {
		w.logger.Error("Get Weapon Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetWeaponResponse](resp), nil
}

func (w *WeaponHandler) GetWeaponList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[myassemblyv1.GetWeaponListResponse], error) {
	resp, err := w.usecase.GetWeaponList(ctx)
	if err != nil {
		w.logger.Error("Get Weapon Error", "detail", err)
		return nil, err
	}

	return connect.NewResponse[myassemblyv1.GetWeaponListResponse](resp), nil
}
