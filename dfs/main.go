package main

import (
	"fmt"
)

type Node struct{
	data int
	neighbours []*Node
	visited bool
	startTime int
	endTime int
}

type Stack struct{
	value []*Node
	size int
}

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

type Tree struct{}

var counter int = 0
var nodesAfterDfs = make([]*Node, 50)

func main(){

	node40 :=  &Node{data:40}
	node10 :=  &Node{data:10}
	node20 :=  &Node{data:20}
	node30 :=  &Node{data:30}
	node60 :=  &Node{data:60}
	node50 :=  &Node{data:50}
	node70 :=  &Node{data:70}

	node40.neighbours = append(node40.neighbours, node10)
	node40.neighbours = append(node40.neighbours, node20)
	node10.neighbours = append(node10.neighbours, node30)
	node20.neighbours = append(node20.neighbours, node10)
	node20.neighbours = append(node20.neighbours, node30)
	node20.neighbours = append(node20.neighbours, node60)
	node20.neighbours = append(node20.neighbours, node50)
	node30.neighbours = append(node30.neighbours, node60)
	node60.neighbours = append(node60.neighbours, node70)
	node50.neighbours = append(node50.neighbours, node70)

	/* fmt.Println(node40.neighbours)
	fmt.Println(node30.neighbours)
	fmt.Println(node10.neighbours) */



	 tree := &Tree{}
	tree.dfsUsingStack(node40) 
}

func (tree *Tree) dfsUsingStack(node *Node){
	dfsStack := &Stack{}
	counter++
	node.startTime = counter
	dfsStack.push(node)
 
	for !dfsStack.isEmpty() {
		poppedElement := dfsStack.pop()
		counter++
		poppedElement.endTime = counter
		poppedElement.visited = true
		fmt.Println(len(poppedElement.neighbours))
		nodesAfterDfs = append(nodesAfterDfs, poppedElement)

		for _ , element := range poppedElement.neighbours {
			if(!element.visited){
				counter++
				element.startTime = counter
				dfsStack.push(element)
			}
		}
	}
}