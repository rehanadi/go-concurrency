package helpers

import "fmt"

// 1.1. Write a function printNumbers that prints numbers from 1 to 10.
func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		println(i)
	}
}

// 1.2. Write a function printLetters that prints letters from 'a' to 'j'.
func PrintLetters() {
	for i := 'a'; i <= 'j'; i++ {
		println(string(i))
	}
}

// 3.1. Create a function produce that sends numbers from 1 to 10 on a channel.
func Produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 3.2. Create a function consume that reads from the channel and prints the numbers.
func Consume(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
