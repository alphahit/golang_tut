package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	var fruitList = []string{"Apple", "Orange", "Bannana"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Printf("Type of fruitlist is %v\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Printf("Fruitlist after appending %v\n", fruitList)

	fruitList = append(fruitList[1:3], "Peach")

	fmt.Println(fruitList)

	highScores := make([]int, 4)
	fmt.Printf("Highscroes : %v\n",highScores)
	
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Printf("Highscroes : %v\n",highScores)

	highScores = append(highScores, 555,666,777)

	fmt.Printf("Highscroes : %v\n",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Printf("Highscroes sorted: %v\n",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//Remove value from slice based on index
	var courses = []string{"ReactJS", "React Native", "JavaScript", "Swift", "Golang"}
	fmt.Println(courses)
	index := 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)


}