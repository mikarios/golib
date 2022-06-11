package contexts

import (
	"context"

	"github.com/mikarios/golib/logger"
)

func Copy(source context.Context, keys ...any) context.Context {
	ctx := context.Background()
	txKey := logger.Settings.TransactionKey

	transactionValue := source.Value(txKey)
	if transactionValue != nil {
		ctx = context.WithValue(ctx, txKey, transactionValue)
	}

	for _, key := range keys {
		ctx = context.WithValue(ctx, key, source.Value(key))
	}

	return ctx
}
