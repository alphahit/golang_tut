package main

import (
	"fmt"
	"math"
)

func findDuplicate(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		if _,exists := m[num]; exists {
			return num
		}
		m[num] = true
	}
	return -1
}


func findDiuplicateMark(nums []int) int {
	for _, num := range nums {
		idx := int(math.Abs(float64(num)))

		if nums[idx] < 0{
			return idx
		}
		nums[idx] = -nums[idx]
	}
	return -1
}

func main(){
	nums := []int{3,1,3,4,2}

	fmt.Println(findDuplicate(nums))

	nums = []int{3,1,3,4,2}

	fmt.Println(findDiuplicateMark(nums))
}