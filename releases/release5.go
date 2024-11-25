package releases

import "fmt"

func Release5() {
	fmt.Println("=== Release 5 ===")

	// 5.1. Create two channels: one for sending even numbers and another for sending odd numbers from 1 to 20.
	even := make(chan int)
	odd := make(chan int)

	// 5.2. Use a single goroutine that sends numbers to these channels based on whether they're even or odd.
	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(even)
		close(odd)
	}()

	// 5.3. In your main function, use the select statement to read from these channels, printing whether an even or odd number was received. Stop the operation after all numbers have been printed.
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
		}
		if even == nil && odd == nil {
			break
		}
	}
}
