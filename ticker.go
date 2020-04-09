package ticker

import (
	"time"
)

type Ticker struct {
	t    *time.Ticker
	stop chan bool
}

func New(duration time.Duration, callback func(time.Time)) *Ticker {
	ticker := &Ticker{
		t:    time.NewTicker(duration),
		stop: make(chan bool),
	}

	go func() {
		for {
			select {
			case t := <-ticker.t.C:
				callback(t)
			case <-ticker.stop:
				return
			}
		}
	}()
	return ticker
}

func (t *Ticker) Stop() {
	t.t.Stop()
	close(t.stop)
}
