package web

import (
	"easy/context"
	"strings"
	"net/http"
	"fmt"
	"easy/web/filter"
	"log"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	DeleteMethod = "DELETE"
)

type controllerFunc func(ctx context.Context) error

type Controller interface {
	Get(context.Context)    error
	Post(context.Context)   error
	Put(context.Context)    error
	Delete(context.Context) error
}

var filters []filter.Filter

type Router struct {
	router map[string]Controller
	filter []filter.Filter
	server *Server
}

func NewRouter(srv *Server) *Router {
	return &Router{
		router: make(map[string]Controller),
		filter: filters,
		server: srv,
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var(
		path   = r.URL.Path
		method = strings.ToUpper(r.Method)
	)
	if _, ok := router.router[path]; !ok {
		http.NotFound(w, r)
		r.Body.Close()
		return
	}
	limiter := router.server.GetInterfaceLimiter()
	ctx := NewContext(w, r, path, limiter)
	controller := router.router[path]
	mapMethodController := map[string]controllerFunc{
		GetMethod:    controller.Get,
		PostMethod:   controller.Post,
		PutMethod:    controller.Put,
		DeleteMethod: controller.Delete,
	}
	if _, ok := mapMethodController[method]; !ok {
		http.NotFound(w, r)
		r.Body.Close()
		return
	}
	filterName, statusCode, err := router.doPreFilters(ctx)
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		log.Println(fmt.Sprintf("call pre filter failed, filter=<%s> statuscode=<%d> errors=<%s>", filterName, statusCode, err.Error()))
		return
	}
	err = mapMethodController[method](ctx)
	if err != nil {
		router.doPostErrorFilters(ctx)
	}
	filterName, statusCode, err = router.doPostFilters(ctx)
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		log.Println(fmt.Sprintf("call pre filter failed, filter=<%s> statuscode=<%d> errors=<%s>", filterName, statusCode, err.Error()))
		return
	}
	r.Body.Close()
}

func (router *Router) doPreFilters(ctx context.Context) (string, int, error) {
	for _, f := range router.filter {
		statusCode, err := f.Pre(ctx)
		if err != nil {
			return f.Name(), statusCode, err
		}
	}
	return "", http.StatusOK, nil
}

func (router *Router) doPostFilters(ctx context.Context) (string, int, error) {
	l := len(router.filter)
	for i := l - 1; i >= 0; i-- {
		f := router.filter[i]
		statusCode, err := f.Post(ctx)
		if err != nil {
			return f.Name(), statusCode, err
		}
	}
	return "", http.StatusOK, nil
}

func (router *Router) doPostErrorFilters(ctx context.Context) {
	l := len(router.filter)
	for i := l - 1; i >= 0; i-- {
		f := router.filter[i]
		f.PostErr(ctx)
	}
}

func (router *Router) RegisterController(path string, controller Controller) {
	router.router[path] = controller
}

func (router *Router) RegisterFilter(f filter.Filter) {
	router.filter = append(router.filter, f)
}

func init() {
	rateLimiterFilter := &RateLimitFilter{}
	filters = append(filters, rateLimiterFilter)
}