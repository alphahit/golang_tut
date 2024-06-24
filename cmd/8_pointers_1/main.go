// Pointers
package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	//var p *int32 = new(int32): Declares a pointer p to an int32 and allocates memory for an int32 value. 
	//Initially, the allocated memory contains the zero value of int32, which is 0.
	var i int32
	fmt.Printf("The value of p points to is: %v", p)
	fmt.Printf("\nThe value of p is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	*p = 10
	//*p = 10: Sets the value at the address p points to, which is now 10.
	fmt.Printf("\n2 - The value of p points to is: %v", p)
	fmt.Printf("\n2 - The value of p is: %v", *p)
	fmt.Printf("\n2 - The value of i is: %v", i)

	p = &i
	//p = &i: Makes p point to the address of i.
	*p = 1
	//*p = 1: Sets the value of i (through p) to 1.
	fmt.Printf("\n3 - The value of p points to is: %v", p)
	fmt.Printf("\n3 - The value of i points to is: %v", &i)
	i = 10
	fmt.Printf("\n4 - The value of p is: %v", *p)
	//Dereferences p to print the value at the address p points to, which is now 10.
	fmt.Printf("\n4 - The value of i is: %v", i)

	*p = 15
	//*p = 15: Sets the value of i (through p) to 15.
	fmt.Printf("\n5 - The value of p is: %v", *p)
	fmt.Printf("\n5 - The value of i is: %v", i)

	var k int32 = 2
	//var k int32 = 2: Declares a variable k of type int32 and initializes it to 2.
	i = k
	//i = k: Sets i to the value of k, which is 2.
	fmt.Printf("\n6 - The value of p is: %v", *p)
	fmt.Printf("\n6 - The value of i is: %v", i)
	fmt.Printf("\n6 - The value of i is: %v", k)

	//slice contains pointers to underlying array
	//slices copying poiters when we do this
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)


	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n7 - The memory location of the thing1 array is: %p",&thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\n The result is: %v", result)

	var thing1Pointer = [5]float64{1,2,3,4,5}
	fmt.Printf("\n7 - The memory location of the thing1Pointer array is: %p",&thing1Pointer)
	var resultPointer [5]float64 = squareByPointer(&thing1Pointer)
	fmt.Printf("\n The result is: %v", resultPointer)
}


func square(thing2[5] float64) [5] float64{
	fmt.Printf("\n The memory location of thing2 array is: %p", &thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}

func squareByPointer(thing2 *[5]float64) [5]float64{
	 // This prints the address of the pointer variable `thing2`
	 fmt.Printf("\n The memory location of thing2 pointer variable is: %p", &thing2)
	 // This prints the address of the array that `thing2` points to
	 fmt.Printf("\n The memory location of the array thing2 points to is: %p", thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}