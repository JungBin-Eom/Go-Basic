package main

import "fmt"

const (
	name = "kim"
	RNN = "800101-1000000"
	job // 따로 초기화 안하면 위의 상수와 같은 값을 가짐
)

func main() {
	fmt.Println(name, RNN, job)
}