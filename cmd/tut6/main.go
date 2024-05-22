//String Runes Bytes

package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "evrēˌdā"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i,v)
	}	
	fmt.Printf("\n The length of myString is %v", len(myString)) 

	var myString1 = []rune("evrēˌdā")
	var indexed1 = myString1[0]
	fmt.Printf("%v, %T\n", indexed1, indexed1)
	for i, v := range myString1{
		fmt.Println(i,v)
	}	
	fmt.Printf("\n The length of myString1 is %v", len(myString1)) 

	var myRune = 'a'
	fmt.Printf("\n myRune is %v", myRune)

	var strSlice = []string{"s","u","b","s","c","r","i","b","e"}
	var catStr = ""
	
	for i := range strSlice{
		catStr += strSlice[i]
		//Everytime a new strign is created, strings in go is immutable
	}
	fmt.Printf("\n%v", catStr)

	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr1 = strBuilder.String()
	fmt.Printf("\n%v", catStr1)
}