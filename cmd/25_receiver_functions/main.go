package main

import "fmt"

//Receiver function is a type of a special kind of function in golang which can accept type not as a parameter but it will accept types and we can call
//the receiver function using a DOT operator on that type

// Syntax :
// func (<structObj><structName>) <funcName> (<Params>) <returnType>{
//	}

type S1 struct {
	name string
	age int
	dob string
}


func main() {
	s1 := S1{
		"Mark",
		22,
		"2010-01-01", //object oriented pattern i.e. calling object public memebers using DOT operator
	}

	fmt.Println(s1)
	name := s1.ReceiverFunc()
	fmt.Println(name)
	s1.ReceiverFuncTwo()
}

func (s1 S1) ReceiverFunc() string {
	return "Name in the struct object is " + s1.name
} 

func (s1 S1) ReceiverFuncTwo() {
	fmt.Println(s1)
} 