# ticker
A thin wrapper for Go’s ticker.

## Installation

``` shell
go get github.com/knwoop/ticker
```

## Example

``` go
package main

import (
	"fmt"
	"time"

	"github/knwoop/ticker"
)

func main() {
	ticker := ticker.New(10*time.Millisecond, func(t time.Time) {
		fmt.Printf("tick at %v", t)
	})

    	time.Sleep(35 * time.Millisecond)
	ticker.Stop()
	time.Sleep(15 * time.Millisecond)
}
```

## License

MIT
