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

//Insert will take a kety and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index:= hash(key)
	h.array[index].insert(key)
}



// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index:= hash(key)
	return h.array[index].search(key)

}

//Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string){
	index:= hash(key)
	h.array[index].delete(key)
}

//Bucket insert
func (b *bucket) insert(k string){
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Already Exists ")
	}
	
}
//Bucket search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//Bucket Delete
func (b *bucket) delete(k string){
	if b.head.key == k {
		b.head =  b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil{
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _,v := range key{
		sum += int(v)
	}
	return sum % ArraySize
}


func Init() *HashTable{
	result := &HashTable{}
	
	for i := range result.array{
		result.array[i] = &bucket{}
	}
	return result
}

func main(){
	hashTable := Init()

	list := []string{
		"ERIC",
		"BANA",
		"KYLE",
		"STAN",
		"RANDY",
		"TOKEN",
	}
	for _,v := range list{
		hashTable.Insert(v)
	}
	fmt.Println(hashTable.Search("STAN"))
	hashTable.Delete("STAN")
	fmt.Println(hashTable.Search("STAN"))
	fmt.Println(hashTable.Search("BANA")) 
	fmt.Println(hashTable)
	
}
