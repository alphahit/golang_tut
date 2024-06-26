package main

import (
	"fmt"
)
var count int
type Node struct {
	Key int
	Left *Node
	Right *Node
}

func  (n *Node) Insert(k int)  {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		}else {
			n.Right.Insert(k)
		}

	}else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		}else{
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool{
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	}else if n.Key > k{
		return n.Left.Search(k)
	}
	return true
} 

func main (){
	tree := &Node{Key : 100}
	fmt.Println(tree)
	tree.Insert(300)
	tree.Insert(50)
	tree.Insert(200)
	fmt.Println(tree)
	fmt.Println(tree.Search(200))
	fmt.Println(count)
	count = 0
	fmt.Println(tree.Search(250))
	fmt.Print(count)
}