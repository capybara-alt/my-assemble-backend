package main

import (
	"context"
	"log/slog"
	"os"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/infrastructure/crawler"
	repo "github.com/capybara-alt/my-assemble/infrastructure/db"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/common"
	"github.com/capybara-alt/my-assemble/usecase/job"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  core.DB_DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: gorm_logger.New(slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo), gorm_logger.Config{
			LogLevel: gorm_logger.Info,
			Colorful: true,
		}),
	})
	if err != nil {
		logger.Error("Failed to connect to database")
	}
	ctx := context.Background()
	ctx = core.SetTx(ctx, db)

	validationSchemaRepo := repo.NewValidationUnitSchema()
	validationSchema := common.NewValidationSchema(validationSchemaRepo, logger)
	validationSchema.Execute(ctx)

	weaponConvert := convert.NewWeaponList(validationSchema.GetWeaponSchema())
	var externalWeaponRepos = []repository.ExternalWeapon{
		crawler.NewCrossWeapon(),
		crawler.NewAmmoWeapon(),
		crawler.NewEnCoralWeapon(),
		crawler.NewExplodeWeapon(),
		crawler.NewExtraWeapon(),
		crawler.NewLauncherWeapon(),
		crawler.NewMissileWeapon(),
		crawler.NewOrbitTaletDroneWeapon(),
		crawler.NewShieldWeapon(),
	}
	weaponRepo := repo.NewWeapon()
	job.NewWeaponJob(weaponRepo, externalWeaponRepos, weaponConvert, logger).Execute(ctx)

	frameConvert := convert.NewFrameList(validationSchema.GetFrameSchema())
	var externalFrameRepos = []repository.ExternalFrame{
		crawler.NewArmsFrame(),
		crawler.NewCoreFrame(),
		crawler.NewHeadFrame(),
		crawler.NewLegsFrame(),
		crawler.NewOtherLegsFrame(),
	}
	frameRepo := repo.NewFrame()
	job.NewFrameJob(frameRepo, externalFrameRepos, frameConvert, logger).Execute(ctx)

	innerUnitConvert := convert.NewInnerUnitsList(validationSchema.GetInnerUnitsSchema())
	var externalInnerUnitRepos = []repository.ExternalInnerUnit{
		crawler.NewBoosterInnerUnit(),
		crawler.NewFcsInnerUnit(),
		crawler.NewGeneratorInnerUnit(),
	}
	innerUnitRepo := repo.NewInnerUnit()
	job.NewInnerUnitJob(innerUnitRepo, externalInnerUnitRepos, innerUnitConvert, logger).Execute(ctx)

	expansionConvert := convert.NewExpansionList(validationSchema.GetExpansionSchema())
	var externalExpansionRepos = []repository.ExternalExpansion{
		crawler.NewExpansion(),
	}
	expansionRepo := repo.NewExpansion()
	job.NewExpansionJob(expansionRepo, externalExpansionRepos, expansionConvert, logger).Execute(ctx)
}
