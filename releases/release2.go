package releases

import (
	"fmt"
	"time"

	h "go-concurrency/helpers"
)

func Release2() {
	// 1.3. Use goroutines to concurrently run both functions.
	fmt.Println("=== Release 2 ===")

	go h.PrintNumbers()
	go h.PrintLetters()

	time.Sleep(5 * time.Second)
}
