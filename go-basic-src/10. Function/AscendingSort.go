package main

import "fmt"

func bubbleSort(nums ...int) (sorted []int) {
	var temp int
	sorted = nums
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	return
}

func inputNums() (numbers []int) {
	var count, number int
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		fmt.Scanln(&number)
		numbers = append(numbers, number)
	}

	return
}

func outputNums(sortednums ...int) {
	for _, num := range sortednums {
		fmt.Printf("%d ", num)
	}
}

func main() {
	nums := inputNums()
	sortednums := bubbleSort(nums...)
	outputNums(sortednums...)
}
