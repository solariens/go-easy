package web

import (
	"net/http"
	"fmt"
	"time"
	"easy/debug"
	"easy/utils"
	"easy/web/filter"
	"golang.org/x/time/rate"
	"context"
	"log"
)

type Server struct {
	handler *Router
	server  *http.Server
	config  FrameworkConf
	debug   *debug.Debug
	limiter map[string]*rate.Limiter
}

func NewServer() *Server {
	server := &http.Server{}
	srv := &Server{server:server}
	srv.handler = NewRouter(srv)
	srv.limiter = make(map[string]*rate.Limiter)
	return srv
}

func (s *Server) GetModuleName() string {
	return s.config.Application
}

func (s *Server) GetInterfaceLimiter() map[string]*rate.Limiter {
	return s.limiter
}

func (s *Server) RegisterController(path string, controller Controller) {
	s.handler.RegisterController(path, controller)
}

func (s *Server) RegisterFilter(f filter.Filter) {
	s.handler.RegisterFilter(f)
}

func (s *Server) Init() {

	s.config = FrameworkConfig

	ip := utils.ParseServerIP(s.config.Server.IP)
	d := debug.NewDebug(s.config.Debug.Enable, ip, s.config.Debug.Port)
	d.Start()

	for _, l := range s.config.Limiter {
		if !l.EnableRateLimit {
			continue
		}
		s.limiter[l.InterfaceName] = rate.NewLimiter(rate.Every(time.Second/time.Duration(l.MaxRequestPerSecond)), int(l.MaxRequestPerSecond))
		log.Println(fmt.Sprintf("request path[%s] add rate limit filter", l.InterfaceName))
	}
}

func (s *Server) Run() {
	ip := utils.ParseServerIP(s.config.Server.IP)
	log.Println(fmt.Sprintf("ip[%s] port[%d] server is starting ...", ip, s.config.Server.Port))

	addr := fmt.Sprintf("%s:%d", ip, s.config.Server.Port)
	s.server.Handler        = s.handler
	s.server.WriteTimeout   = time.Duration(s.config.Server.WriteTimeout) * time.Millisecond
	s.server.ReadTimeout    = time.Duration(s.config.Server.ReadTimeout) * time.Millisecond
	s.server.IdleTimeout    = time.Duration(s.config.Server.IdleTimeout) * time.Millisecond
	s.server.MaxHeaderBytes = s.config.Server.MaxHeaderSize
	s.server.Addr = addr
	log.Fatal(s.server.ListenAndServe())
}

func (s *Server) Close() {
	logger.Close()
	s.server.Shutdown(context.Background())
}