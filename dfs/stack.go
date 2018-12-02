package main

import (
	"fmt"
)

type Stack struct{
	value []*Node
	size int
}

/* func main(){
	s := &Stack{}

	s.push(1)
	s.push(2)
	s.push(3)

	s.printStack()

	fmt.Println()

	fmt.Println(s.pop())

	s.printStack()

	fmt.Println()

	fmt.Println(s.peek())

	s.printStack()
} */

func (s *Stack) push(node *Node){
	s.value = append(s.value,node)
	s.size++
}

func (s *Stack) pop()(*Node){
	if s.size == 0 {
		return nil
	}
	s.size--
	return s.value[s.size]
}

func (s *Stack) peek()(*Node){
	if s.size == 0 {
		return nil
	}
	return s.value[s.size-1]
}

func (s *Stack) printStack(){
	for i := 0 ; i < s.size ; i++ {
		fmt.Print(s.value[i])
	}
}

func (s *Stack) isEmpty()(bool){
	return s.size == 0
}