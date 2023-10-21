package interceptor

import (
	"context"
	"os"

	"log/slog"

	"connectrpc.com/connect"
	"github.com/capybara-alt/my-assemble/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func NewInitConnectionInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (resp connect.AnyResponse, err error) {
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
				return nil, err
			}

			err = db.Transaction(func(tx *gorm.DB) error {
				resp, err = next(core.SetTx(ctx, tx), req)
				if err != nil {
					return err
				}
				return nil
			})

			return resp, err
		})
	}
}
