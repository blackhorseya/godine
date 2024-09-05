package contextx

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.NewExample())
}

// ContextKey is a type for context key.
type ContextKey string

const (
	// KeyContextx is the key of contextx.
	// Deprecated:
	KeyContextx ContextKey = "contextx"
)

// Contextx extends google's context to support logging methods.
type Contextx struct {
	context.Context
	*zap.Logger
}

// WithContextx returns a copy of parent in which the context is set to c.
func WithContextx(c context.Context) Contextx {
	return Contextx{
		Context: c,
		Logger:  ctxzap.Extract(c),
	}
}

// WithContextLegacy returns a copy of parent in which the context is set to ctx.
// Deprecated: Use WithContextx instead.
func WithContextLegacy(ctx context.Context) Contextx {
	return Contextx{
		Context: ctx,
		Logger:  zap.L(),
	}
}

// Background returns a non-nil, empty Contextx. It is never canceled, has no values, and has no deadline.
// Deprecated: Use WithContextx instead.
func Background() Contextx {
	return Contextx{
		Context: context.Background(),
		Logger:  zap.L(),
	}
}
