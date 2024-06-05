package main

import "fmt"

//ArraySize is the size of hash table array
const ArraySize = 7

//HashTable will holf an array
type HashTable struct {
	array [ArraySize]*bucket
}

//bucket is a linked list in each slot of 
type bucket struct {
	head *bucketNode
}


type bucketNode struct {
	key string
	next *bucketNode
}

func Init() *HashTable{
	result := &HashTable{}
	
	for i := range result.array{
		result.array[i] = &bucket{}
	}
	return result
}

func main(){
	testHashTable := &HashTable{}
	fmt.Println(testHashTable)

}
