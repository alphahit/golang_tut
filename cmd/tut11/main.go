//Channels

//In high level we can understand chanels as goroutines which pass around information

//They hold data : Integer, Slice
//They are thread safe, i.e. we avoid data races when we are reading and writing from memory
//We can listen and remove when a data is added or removed from the channek and we can block code execution until one of these events happens

package main

import "fmt"

func main(){
	var c = make(chan int)
	go process(c)
	for i:= range c{
		fmt.Println(i)
	}
}

func process(c chan int){
	defer close(c)
	for i:=0; i<5; i++{
		c <- i
	}
}