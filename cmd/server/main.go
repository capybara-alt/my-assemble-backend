package main

import (
	"log/slog"
	"net/http"
	"os"

	"buf.build/gen/go/capybara/my-assemble/connectrpc/go/myassembly/v1/myassemblyv1connect"
	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/handler"
	repo "github.com/capybara-alt/my-assemble/infrastructure/db"
	"github.com/capybara-alt/my-assemble/interceptor"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	interceptors := connect.WithInterceptors(interceptor.NewInitConnectionInterceptor())
	expansionRepo := repo.NewExpansion()
	frameRepo := repo.NewFrame()
	innerUnitRepo := repo.NewInnerUnit()
	weaponRepo := repo.NewWeapon()
	validationUnitSchemaRepo := repo.NewValidationUnitSchema()
	assemblyRepo := repo.NewAssembly()
	userRepo := repo.NewUser()
	mux.Handle(myassemblyv1connect.NewExpansionServiceHandler(
		handler.NewExpansionHandler(expansionRepo, logger),
		interceptors,
	))
	mux.Handle(myassemblyv1connect.NewFrameServiceHandler(
		handler.NewFrameHandler(frameRepo),
		interceptors,
	))
	mux.Handle(myassemblyv1connect.NewInnerUnitServiceHandler(
		handler.NewInnerUnitHandler(innerUnitRepo),
		interceptors,
	))
	mux.Handle(myassemblyv1connect.NewWeaponServiceHandler(
		handler.NewWeaponHandler(weaponRepo, logger),
		interceptors,
	))
	mux.Handle(myassemblyv1connect.NewValidationUnitSchemaServiceHandler(
		handler.NewValidationUnitSchemaHandler(validationUnitSchemaRepo, logger),
		interceptors,
	))
	mux.Handle(myassemblyv1connect.NewAssemblyServiceHandler(
		handler.NewAssemblyHandler(assemblyRepo, logger),
		interceptors,
	))
	mux.Handle(myassemblyv1connect.NewUserServiceHandler(
		handler.NewUserHandler(userRepo, logger),
		interceptors,
	))

	logger.Info("Start server")
	err := http.ListenAndServe("0.0.0.0:8080", h2c.NewHandler(mux, &http2.Server{}))
	logger.Error("Failed to start server", "detail", err)
}
