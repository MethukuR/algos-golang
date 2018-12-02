package main

import (
	"fmt"
	"math"
)

type Stack struct{
	value []int
	size int
}

func main(){
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


	fmt.Println(s.pop())

	fmt.Println(s.pop())

	fmt.Println(s.pop())

	s.printStack()

	fmt.Println(s.isEmpty())



}

func (s *Stack) push(data int){
	s.value = append(s.value,data)
	s.size++
}

func (s *Stack) pop()(int){
	if s.size == 0 {
		return math.MinInt64
	}
	s.size--
	return s.value[s.size]
}

func (s *Stack) peek()(int){
	if s.size == 0 {
		return math.MinInt64
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