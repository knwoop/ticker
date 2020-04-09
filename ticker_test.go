package ticker_test

import (
	"fmt"
	"time"

	"github.com/knwoop/ticker"
)

func ExampleTicker() {
	ticker := ticker.New(10*time.Millisecond, func(_ time.Time) {
		fmt.Println("ticked")
	})

	time.Sleep(35 * time.Millisecond)
	ticker.Stop()
	time.Sleep(15 * time.Millisecond)
}
