package main

import "fmt"

func main() {
	var num1, num2, num3 int
	
	fmt.Scanln(&num1, &num2, &num3)
	
	// go에서는 자료형 크기를 자주 명시해준다.
	num11 := float32(num1)
	num22 := uint(num2)
	num33 := int64(num3)
	
	fmt.Printf("%T, %f\n%T, %d\n%T, %d\n", num11, num11, num22, num22, num33, num33)
	// %T는 타입 출력
}