package debug

import (
	_ "net/http/pprof"
	"log"
	"net/http"
	"fmt"
)

type Debug struct {
	enable bool
	port   int
	ip     string
}

func NewDebug(enable bool, ip string, port int) *Debug {
	return &Debug{
		enable: enable,
		port:   port,
		ip:     ip,
	}
}

func (d *Debug) Start() {
	if !d.enable {
		return
	}
	go func() {
		log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", d.ip, d.port), nil))
	}()
}
