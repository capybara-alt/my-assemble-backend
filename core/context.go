package core

import (
	"context"

	"gorm.io/gorm"
)

type ContextKey string

const tx_context_key ContextKey = "dbc"

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, tx_context_key, tx)
}

func GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(tx_context_key).(*gorm.DB); ok {
		return tx
	}

	return nil
}
