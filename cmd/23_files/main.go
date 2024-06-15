package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Prateek's go journey"

	file, err := os.Create("./mygofile.txt")
	
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Lenght is: ", length)

	defer file.Close()
	readFile("./mygofile.txt")
}


func readFile(fileName string) {
	//Whenever you read the file it's not being read into the sting format
	//It is alway rea into byte format
	dataByte, err := os.ReadFile(fileName)

	checkNilErr(err)
	

	fmt.Println("data that is read from the file is: \n", dataByte)
	fmt.Println("text read from the file is: \n", string(dataByte))
}

func checkNilErr(err error){
	if err != nil {
		//panic will just shutown the program an will show you what the error you are facing
		panic(err)
	}
}