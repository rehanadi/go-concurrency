package releases

import "fmt"

func Release6() {
	fmt.Printf("\n\n=== Release 6 ===\n")

	// 6.1. Implement a scenario where you can pass errors between goroutines using channels. For example, when trying to send a number greater than 20 to the channels in Release 5, send an error to an error channel.
	// 6.2. Handle this error in the main function by printing it out.

	even := make(chan int)
	odd := make(chan int)
	over := make(chan error)

	go func() {
		for i := 1; i <= 22; i++ {
			if i > 20 {
				over <- fmt.Errorf("number %d is greater than 20", i)
				continue
			}
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(even)
		close(odd)
		close(over)
	}()

	for {
		select {
		case num, ok := <-even:
			if !ok {
				even = nil
			} else {
				fmt.Println("Received an even number", num)
			}
		case num, ok := <-odd:
			if !ok {
				odd = nil
			} else {
				fmt.Println("Received an odd number", num)
			}
		case err, ok := <-over:
			if !ok {
				over = nil
			} else {
				fmt.Println("Error:", err)
			}
		}
		if even == nil && odd == nil && over == nil {
			break
		}
	}
}
