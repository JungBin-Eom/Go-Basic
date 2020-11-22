package main

import "fmt"

func main() {
	var RNNf, RNNt int
	var name string
	var height float32
	
	fmt.Scanf("%d-%d", &RNNf, &RNNt)
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &height)
	
	fmt.Printf("주민등록번호 앞자리는 %d, 뒷자리는 %d, 이름은 %s입니다.\n", RNNf, RNNt, name)
	fmt.Printf("그리고 키는 %.2f입니다.", height)
}