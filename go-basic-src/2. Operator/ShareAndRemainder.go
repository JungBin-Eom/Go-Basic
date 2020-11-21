package main

import "fmt"

func main() {
	var num1, num2, result1, result2 int
	
	fmt.Scanln(&num1, &num2)
	
	result1 = num1/num2
	result2 = num1%num2
	
	fmt.Printf("몫 : %d, 나머지 : %d", result1, result2)
}