package main

import "fmt"

func main() {
	var i, j, length int
	//i,j는 두 개의 반복문에 쓰일 변수
	
	fmt.Scanln(&length)
	
	for i = 0; i < length; i++ {
		for j = 0; j < i; j++ {
			fmt.Print("o ")
		}
		fmt.Println("* ")
	}
}