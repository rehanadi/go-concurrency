package releases

import (
	"fmt"
)

// 3.1. Create a function produce that sends numbers from 1 to 10 on a channel.
func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 3.2. Create a function consume that reads from the channel and prints the numbers.
func consume(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func Release3() {
	// 3.3. Use goroutines to concurrently run both functions and demonstrate communication between them using the channel.
	fmt.Println("=== Release 3 ===")

	ch := make(chan int)

	go produce(ch)

	consume(ch)
}
