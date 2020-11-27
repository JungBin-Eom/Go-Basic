package main

import "fmt"

func inputNums() (scores []int) {
	var score int
	scores = make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		fmt.Scanln(&score)
		scores = append(scores, score)
	}

	return
}

func calExam(arr ...int) (total int, over90 int, under70 int) {
	for _, score := range arr {
		total += score
		if score >= 90 {
			over90 += 1
		} else if score < 70 {
			under70 += 1
		}
	}

	return
}

func printResult(total int, over90 int, under70 int) {
	var result bool = true

	if total < 400 {
		result = false
		fmt.Println("총점이 400점 미만입니다.")
	}
	if over90 < 2 {
		result = false
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
	}
	if under70 != 0 {
		result = false
		fmt.Println("70점 미만 과목이 있습니다.")
	}

	if result == true {
		fmt.Println("아이패드를 살 수 있습니다.")
	} else {
		fmt.Println("아이패드를 살 수 없습니다.")
	}
}

func main() {
	scores := inputNums()
	printResult(calExam(scores...))
}
