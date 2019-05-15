package context

import (
	"context"
	"time"
)

type Context interface {
	context.Context
	WithTimeout(time.Duration) (context.Context, context.CancelFunc)
}
