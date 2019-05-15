package web

import (
	"easy/context"
	"net/http"
	"easy/web/filter"
	stdCtx "context"
)

const (
	RateLimitFilterName = "RateLimit"
)

type RateLimitFilter struct {
	filter.BaseFilter
}

func (r *RateLimitFilter) Name() string {
	return RateLimitFilterName
}

func (r *RateLimitFilter) Pre(ctx context.Context) (statusCode int, err error) {
	path := ctx.(*HttpContext).InterfaceName
	limiter := ctx.(*HttpContext).PathLimiter(path)
	if limiter == nil {
		return http.StatusOK, nil
	}
	err = limiter.Wait(stdCtx.Background())
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return r.BaseFilter.Pre(ctx)
}
