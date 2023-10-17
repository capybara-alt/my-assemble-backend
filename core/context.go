package core

import (
	"context"

	"gorm.io/gorm"
)

type ContextKey string

const txContextKey ContextKey = "dbc"

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txContextKey, tx)
}

func GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txContextKey).(*gorm.DB); ok {
		return tx
	}

	return nil
}
