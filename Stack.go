package main

import "fmt"

type Node struct {
	value int
}

type Stack []*Node

func (s *Stack) Push(elem *Node) {
	*s = append(*s, elem)
}
func (s *Stack) PrintAll() string {
	var elems string = ""
	if len(*s) == 0 {
		fmt.Println("stack is empty")
		return "stack is empty"
	} else {
		for _, v := range *s {
			fmt.Printf("%+v", v.value)
			fmt.Println()
		}

	}

	return elems
}
func (s *Stack) Pop() *Node {
	var lastIndex = len(*s) - 1
	var lastElement = (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return lastElement
}

func main() {
	var stack Stack
	stack.Push(&Node{1})

	stack.Push(&Node{2})
	stack.PrintAll()
	//fmt.Println()
	//fmt.Printf("%+v", stack.Pop())
	fmt.Printf("%+v", stack.Pop())
	fmt.Printf("%+v", stack.Pop())
	fmt.Println()
	stack.PrintAll()
}
