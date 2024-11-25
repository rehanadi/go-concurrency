package releases

import (
	"fmt"
	"time"

	h "go-concurrency/helpers"
)

func Release1() {
	// 1.3. Use goroutines to concurrently run both functions.
	fmt.Printf("\n=== Release 1 ===\n")

	go h.PrintNumbers()
	go h.PrintLetters()

	time.Sleep(2 * time.Second)
}
