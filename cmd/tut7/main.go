//Structs and Interfaces

package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
	int
}

type owner struct{
	name string
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
  return e.gallons*e.mpg 
}
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh*e.kwh 
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need To Fuel up First!")
	}
}

func main(){
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 10}


	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)

	//Not Reusable
	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	var engineCheck gasEngine = gasEngine{10,20, owner{"James"}, 15}
	var engineCheck2 electricEngine = electricEngine{20,30}
	canMakeIt(engineCheck, 250)
	canMakeIt(engineCheck2, 60)
}