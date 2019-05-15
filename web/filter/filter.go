package filter

import (
	"easy/context"
	"net/http"
)

type Filter interface {
	Name() string
	Pre(ctx context.Context) (statusCode int, err error)
	Post(ctx context.Context) (statusCode int, err error)
	PostErr(ctx context.Context)
}

type BaseFilter struct{}

const (
	BaseFilterName = "BaseFilter"
)

func (b *BaseFilter) Name() string {
	return BaseFilterName
}

func (b *BaseFilter) Pre(ctx context.Context) (statusCode int, err error) {
	return http.StatusOK, nil
}

func (b *BaseFilter) Post(ctx context.Context) (statusCode int, err error) {
	return http.StatusOK, nil
}

func (b *BaseFilter) PostErr(ctx context.Context) {

}