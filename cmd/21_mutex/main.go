package main

import (
	"fmt"
	"sync"
)

//A mutex or mutual exclusion is a way for us to esure that at one place that
//like one that we have a spot in memory that spot n memory is only going to be written to by one thing

type SafeCounter struct {
	mu sync.Mutex //sync.Mutex type variable for mutual exclusion (mutex).
	NumMap map[string]int
}

func (s *SafeCounter) Add(num int){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NumMap["key"] = num
}

func main(){

	//Maps
	languages := make(map[string]string) //[key]value 
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages: ", languages)

	//Loops in maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	//Mutex
	s := SafeCounter{NumMap: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()
	fmt.Print(s.NumMap["key"])
}