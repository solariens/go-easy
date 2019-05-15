package easy

import (
	"easy/web"
	"easy/web/filter"
)

var HTTP = web.NewServer()

func RegisterController(path string, controller web.Controller) {
	HTTP.RegisterController(path, controller)
}

func RegisterFilter(filter filter.Filter) {
	HTTP.RegisterFilter(filter)
}

func Run() {
	Init()
	HTTP.Init()
	frameworkIns.run()
}
