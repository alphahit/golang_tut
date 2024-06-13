// Go routine is going to spin out a separate process from your original one
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Wg's shared memory reference
func main() { 
	var wg sync.WaitGroup
	//wg counter goes up
	wg.Add(1)
	go func() {
		//Decrement the wg counter once the function is done
		defer wg.Done()
		for i:= 0; i< 100; i++ {
			fmt.Println("IDX from first func:",i)
			time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
		}
	}()
	
	//wg counter goes up
	wg.Add(1)
	go func(){
		//Decrement the wg counter once the function is done
		defer wg.Done()
		for i:= 0; i< 100; i++ {
			fmt.Println("IDX from second func:",i)
			time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Done ")

	//We can run the above two things in parallel but we can get data in and out of goroutines and the answer to that is channels
	//Channels allows us to pass messages back and forth between different goroutines
	 c := make(chan int)
	 go func(){
		sum := 0
		for i:= 0; i< 100; i++ {
			fmt.Println("IDX from first func:",i)
			sum += i
		}
		c <- sum
	 }()
	 output := <-c //This will block execution until it gets something from the channel
	 fmt.Println("Output: ", output)
	 //Channels are ways to pass messages back and forth between different threads and different goroutines
	 //It makes it easy to synchronize them and to keep data going back and forth
	
}