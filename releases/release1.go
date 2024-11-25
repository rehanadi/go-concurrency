package releases

import (
	"fmt"
	"time"

	h "go-concurrency/helpers"
)

func Release1() {
	// 1.3. Use goroutines to concurrently run both functions.
	fmt.Println("=== Release 1 ===")

	go h.PrintNumbers()
	go h.PrintLetters()

	time.Sleep(2 * time.Second)
}
