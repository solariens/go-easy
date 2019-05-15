package easy

import (
	"easy/signal"
	"log"
)

var frameworkIns = &framework{
	signalTrigger: signal.NewSignalTrigger(),
	quit:          make(chan struct{}),
}

type framework struct {
	signalTrigger *signal.SignalTrigger
	quit          chan struct{}
}

func (f *framework) run() {
	f.signalTrigger.Run()

	HTTP.Run()

	<- f.quit
}

func (f *framework) exit() {
	log.Println("server is closing ...")
	HTTP.Close()
	f.signalTrigger.Close()
	close(f.quit)
	log.Println("server is closed !!!")
}
