package main

import (
	"fmt"
	"log"
	"runtime"
)

//All of your packages are named with one word & all lowercase and if you deal with any  collisions anywhere you can change your import names from there
func main() {

//Variable names should be camelcase		
 testVar := "test" //Good
 test_var := "test" //Bad

 fmt.Print(testVar)
 fmt.Print(test_var)

 //Switch statements and if statements can take initializations as part of their syntax
 if err := foo(); err != nil {
	log.Print(err)
 }

 switch os:= runtime.GOOS; os{
 case "darwin":
	fmt.Println("OS X")
 case "linux":
	fmt.Println("Linux")
 case "windows":
	fmt.Println("Windows")
 default:
	fmt.Printf("%s.\n", os)
 }

 //Reassignments using colon  equals operator
 //With colon equals operator we can't redclare something 
 c := 0
//  c := 1
fmt.Println(c)


//Here I am not redecalring the err variable
//Go understands that if I ahve atlease one new variable on the left
//Then I can declare any number of old variables on the right and reassign them 
a,err := foo1()
 if err != nil {
	log.Print(err)
 }
	fmt.Println(a)
 

 b,err := foo1()
 if err != nil {
	log.Print(err)
 }
	fmt.Println(b)
 

}

func foo() error {
	return fmt.Errorf("foo error")
}

func foo1() (int,error) {
	return 0,nil
}

