package timer

import (
	"fmt"
	"time"
)

// usage `defer Timer("function main took")()`
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
