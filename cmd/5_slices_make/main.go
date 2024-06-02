package main

import (
	"fmt"
	"time"
)

//Slices are of flexible length. In go we mostly will work on slices.
//A slice is built on top of an array
//Whenever you create a slice in memory you are actually creating
//what's called a backing array and the array is what contains the actual data and the slice
//just contains a reference to the data and the underlying array and for this reason the slice is
//known as reference types.
//Slices are reference to a contiguous portion of an underlying array

//Slices can be created with the built-in make function.
//Whenver you create a slice you are creating an underlying array in memory.
func main() {
	var n int = 1000000

	var testSlice = []int{}

	//Make accepts three paramters 
	//1st parameter- type : refers to the type of data that you want to store in a slice
	//2nd paramter- length :  it is the actual length of the slice
	//3th parameter- capacity : refers to the maximum amount of items that can be stored in the underlying array
	mySlice := make([]string, 3, 8)
	fmt.Printf("Length is %d\n capacity is : %d\n",len(mySlice), cap(mySlice))

	mySlice1 := make([]string, 3)
	fmt.Printf("Length is %d\n capacity is : %d\n",len(mySlice1), cap(mySlice1))
	fmt.Println(mySlice1)

	//Another way to create a slice
	mySlice2 := []string{"a", "b", "c"}
	fmt.Printf("Length is %d\n capacity is : %d\n",len(mySlice2), cap(mySlice2))
	fmt.Println(mySlice2)

	mySlice3 := []int{1, 2, 3,8, 4,5,9}
	fmt.Printf("Length is %d\n capacity is : %d\n",len(mySlice3), cap(mySlice3))
	fmt.Println("Second element of mySlice3:", mySlice3[1])
	fmt.Println("Slice from index 1 (inclusive) to 3 (exclusive):", mySlice3[1:3])
	fmt.Println("Slice from beginning to index 3 (exclusive):", mySlice3[:3])
	fmt.Println("Slice from index 1 (inclusive) to the end:", mySlice3[1:])
	mySlice3[0] = 33
	fmt.Println("Changed Slice :=>",mySlice3)
	newSlice :=  mySlice3[1:]
	fmt.Println("newSlice is :=>",newSlice)


	var testSlice2 = make([]int, 0, n)
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	
	var t0 = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
