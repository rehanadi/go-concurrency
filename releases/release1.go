package releases

import (
	"fmt"
	"time"
)

// 1.1. Write a function printNumbers that prints numbers from 1 to 10.
func printNumbers() {
	for i := 1; i <= 10; i++ {
		println(i)
	}
}

// 1.2. Write a function printLetters that prints letters from 'a' to 'j'.
func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		println(string(i))
	}
}

func Release1() {
	// 1.3. Use goroutines to concurrently run both functions.
	fmt.Println("=== Release 1 ===")

	go printNumbers()
	go printLetters()

	time.Sleep(5 * time.Second)
}
