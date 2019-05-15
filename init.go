package easy

import (
	"easy/web"
	"syscall"
)

func Init() {
	frameworkIns.signalTrigger.Bind(frameworkIns.exit, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)
	web.InitConfig()
	web.InitLogger()
}
