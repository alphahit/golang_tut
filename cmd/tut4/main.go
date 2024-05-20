//Map is a set of key valu pairs where you can llok up a value by it's key

package main

import "fmt"
func main() {
	var myMap map[string]uint8 =  make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"Adam":23, "Sarah":23, "Alice":25, "Bob":53}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])

	delete(myMap2,"Sarah")
	var age,ok = myMap2["Adam"]
	if ok{
		fmt.Printf("Age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	for name := range myMap2{
		fmt.Printf("Name: %v\n", name)
	}
	for name, age := range myMap2{
		fmt.Printf("Name: %v  Age: %v\n", name,age)
	}

	intArr := [3]int32{1,2,3}
	for i,v := range intArr{
		fmt.Printf("Index: %v\n Value: %v",i,v)
	}

	i := 0
	for i<10{
		fmt.Println(i)
		i = i+1
	}

	j := 0
	for j<10{
		if j >= 10{
			break
		}
		fmt.Println(j)
		j = j+1
	}

	for l:=0; l<=10; l++{
		fmt.Println(l)
	}
}