//Sum of values of integer slice

package main

import (
	"fmt"
)

func main(){
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumFloat32Slice(float32Slice))

	fmt.Println(sumSlice(intSlice))
	fmt.Println(sumSlice((float32Slice)))

	var intEmpty = []int{}
	var float32Empty = []float32{}

	fmt.Println(isEmpty(intEmpty))
	fmt.Println(isEmpty(float32Empty))

}

func sumIntSlice(intSlice []int) int {
	var sum int
	for _, i := range intSlice {
		sum += i
	}
	return sum
}


func sumFloat32Slice(float32Slice []float32) float32 {
	var sum float32
	for _, i := range float32Slice {
		sum += i
	}
	return sum
}

//Generics Here as a way of allowing function to receive additional parameters within square brackets


func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _,v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}