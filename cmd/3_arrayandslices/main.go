package main

import "fmt"

//Fixed Length
//Same type
func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	arr := arrayTest()
	fmt.Println(arr)

	arrSlice, arrSlice2 := sliceTest()
	fmt.Println(arrSlice, arrSlice2)
}


func arrayTest() [3]int32{
	intArr := [3]int32{1,2,3}
	return intArr
}

func sliceTest() ([]int32, []int32){
	intSlices := []int32{4,5,6}
	fmt.Printf("Length %v Capacity %v", len(intSlices),cap(intSlices))
	intSlices = append(intSlices, 7)
	fmt.Printf("Length %v Capacity %v", len(intSlices),cap(intSlices))
	intSlices2 := []int32{7,8,9}
	fmt.Println("")
	intSlices = append(intSlices, intSlices2...)

	var instSlice3 []int32 = make([]int32, 3, 8)
	return intSlices, instSlice3
}