package client

import (
	"context"
)

type ctxKey byte

const (
	ctxUIDKey ctxKey = iota
)

// ContextWithUID returns a new Context that carries value u.
func ContextWithUID(ctx context.Context, u int64) context.Context {
	return context.WithValue(ctx, ctxUIDKey, u)
}

// UIDFromContext returns the User value stored in ctx, if any.
func UIDFromContext(ctx context.Context) (int64, bool) {
	u, ok := ctx.Value(ctxUIDKey).(int64)
	return u, ok
}
