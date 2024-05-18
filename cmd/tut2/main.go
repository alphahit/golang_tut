package main

import (
	"errors"
	"fmt"
)

func main(){
	var printSome string = "Hello world!"
	printMe(printSome)
	
	if 1==1 && 2==2 {
		fmt.Println("Passed the check")
	}
 

	numerator, denominator := 5,2
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)
	var result1, remainder1, err  =   intDivisionWithRemainder(numerator, denominator)
	
	
	
	if err != nil {
		fmt.Println(err.Error())
	}else if remainder1 == 0 {
		fmt.Printf("The result of division is: %v", result)
	}else{
		fmt.Printf("The result of division is %v and the remainder is %v", result1, remainder1)
	}



	switch{
	case err != nil: fmt.Println(err.Error())
	case remainder1 == 0: fmt.Printf("Switch : e result of division is %v", result1)
	default : fmt.Printf("Switch : The result of division is %v and the remainder is %v", result1, remainder1)
	}
	
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(a int, b int) int {
	var result int = a/b
	return result
}

func intDivisionWithRemainder(a int, b int) (int,int, error) {
	var err error
	if b == 0{
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = a/b
	var remainder int = a % b
	return result, remainder, err
}