package main

import "fmt"

func add(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	var num1, num2 int
	channel := make(chan int)

	fmt.Scanln(&num1, &num2)

	go add(num1, num2, channel)

	fmt.Println(<-channel)
}
