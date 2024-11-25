package releases

import (
	"fmt"
	h "go-concurrency/helpers"
)

func Release4() {
	// 4.1. Modify the channel in Task 3 to have a buffer of size 5.
	fmt.Printf("\n\n=== Release 4 ===\n")

	ch := make(chan int, 5)

	go h.Produce(ch)

	h.Consume(ch)

	// 4.2. Discuss or note down the behavior difference when using a buffered channel versus an unbuffered one.
	/*
		When using a buffered channel, the producer can send multiple values to the channel without blocking, as long as the buffer is not full.
		When using an unbuffered channel, the producer will block until the consumer is ready to receive the value.
	*/
}
