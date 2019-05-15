package web

import (
	"easy/trace"
	"net/http"
	"time"
	"context"
	"golang.org/x/time/rate"
)

type HttpContext struct {
	Request       *http.Request
	Response      http.ResponseWriter
	Trace         *trace.Trace
	ModuleName    string
	InterfaceName string
	Logger        *Logger
	Metadata      MD
	Limiter       map[string]*rate.Limiter
	ctx           context.Context
}

func NewContext(w http.ResponseWriter, r *http.Request, interfaceName string, limiter map[string]*rate.Limiter) *HttpContext {
	t := trace.NewTrace()
	t.FromHTTPHeader(r.Header)
	t.Init(FrameworkConfig.Application, interfaceName)
	logger.SetRecordPrefix(t)
	return &HttpContext{
		Request:       r,
		Response:      w,
		ModuleName:    FrameworkConfig.Application,
		InterfaceName: interfaceName,
		Trace:         t,
		Logger:        logger,
		Metadata:      NewMetadata(),
		Limiter:       limiter,
		ctx:           context.Background(),
	}
}

func (w *HttpContext) PathLimiter(path string) *rate.Limiter {
	var limiter *rate.Limiter
	if l, ok := w.Limiter[path]; ok {
		limiter = l
	}
	return limiter
}

func (w *HttpContext) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(w.ctx, timeout)
	return ctx, cancel
}

func (w *HttpContext) Deadline() (deadline time.Time, ok bool) {
	return w.ctx.Deadline()
}

func (w *HttpContext) Done() <-chan struct{} {
	return w.ctx.Done()
}

func (w *HttpContext) Err() error {
	return w.ctx.Err()
}

func (w *HttpContext) Value(key interface{}) interface{} {
	return w.ctx.Value(key)
}
