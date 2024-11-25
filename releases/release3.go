package releases

import (
	"fmt"
	h "go-concurrency/helpers"
)

func Release3() {
	// 3.3. Use goroutines to concurrently run both functions and demonstrate communication between them using the channel.
	fmt.Printf("\n\n=== Release 3 ===\n")

	ch := make(chan int)

	go h.Produce(ch)

	h.Consume(ch)
}
