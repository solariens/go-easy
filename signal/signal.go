package signal

import (
	"os"
	"os/signal"
)

type SignalTrigger struct {
	signalCh chan os.Signal
	trigger  map[os.Signal]func()
	quit     chan struct{}
}

func NewSignalTrigger() *SignalTrigger {
	trigger := &SignalTrigger{
		signalCh: make(chan os.Signal, 1),
		trigger:  make(map[os.Signal]func()),
		quit:     make(chan struct{}),
	}

	go func() {
		for {
			select {
			case sig, ok := <-trigger.signalCh:
				if ok {
					if fn, ok := trigger.trigger[sig]; ok {
						fn()
					}
				}
			case <-trigger.quit:
				return
			}
		}
	}()

	return trigger
}

func (s *SignalTrigger) Bind(fn func(), signals ...os.Signal) {
	for _, sig := range signals {
		s.trigger[sig] = fn
	}
}

func (s *SignalTrigger) Run() {
	var signals []os.Signal
	for sig := range s.trigger {
		signals = append(signals, sig)
	}
	signal.Notify(s.signalCh, signals...)
}

func (s *SignalTrigger) Close() {
	signal.Stop(s.signalCh)
	close(s.quit)
}
