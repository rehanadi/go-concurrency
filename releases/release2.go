package releases

import (
	"fmt"
	"sync"

	h "go-concurrency/helpers"
)

func Release2() {
	// 2.1. Recognize the problem when the main function doesn't wait for the goroutines to finish.
	// 2.2. Use the sync.WaitGroup to make sure your main function waits for the goroutines to complete before exiting.
	fmt.Printf("\n\n=== Release 2 ===\n")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		h.PrintNumbers()
	}()

	go func() {
		defer wg.Done()
		h.PrintLetters()
	}()

	wg.Wait()
}
