package helper

import (
	"context"
	"fmt"
)

type ctxKey string

func CtxNewValue(ctx context.Context, key ctxKey, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func CtxGetValue(ctx context.Context, key ctxKey) string {
	result := ctx.Value(key)
	if result != nil {
		return result.(string)
	}
	return fmt.Sprintf("warning. value %s not set in context", key)
}
