package main

import (
	"fmt"
)

//Stuct Methods
//Declare a struct and declare special methods that will be directly attached to that struct
type Showcase struct {
	Name string
}

func (s *Showcase) SetName(name string) { 
	s.Name = name
}

func main() {
	type TestStruct struct {
		Name string
	}
	
	// All the new keyword does is basically create a instance of the type you want.
	// However instead of returning the plain declaration of the type,
	// it references it and return the acutal memory address of that type in the program process heap.
	//New wil initialize and allocate the memory for a new variable with all values initialized to zero
	t := new(TestStruct)  //type *TestStruct
	var v TestStruct  //type TestStruct

	fmt.Println(t)
	fmt.Println(v)

	//Make is only used for a slice, channel or a map
	//Make is actually initializing some underlying data structures taht are required
	slice := make([]int, 10)
	fmt.Println(slice)

	//Underscore Operator
	//It allow us to safely ignore the return type from functiom
	//Syntax error in go if we initialize a variable and do not use it 
	i,_ := foo()
	fmt.Println(i)

	var showCase Showcase
	showCase.SetName("Test Show")
	fmt.Println(showCase.Name)
}

func foo() (int, error) {
	return 1, nil
}