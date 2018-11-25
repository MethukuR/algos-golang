package main

import (
	"fmt"
)

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func main(){
	ll := &LinkedList{}
	fmt.Println("Pushed 1, 2 and 3");
	ll.push(1)
	ll.push(2)
	ll.push(3)

	size := ll.size()
	fmt.Println("Size of list", size)

	fmt.Println("Insert 5 at first  ")
	ll.insertFirst(5)
	fmt.Println("First element of list ",ll.first().data)
	fmt.Println("Size of list", ll.size())

	fmt.Println("Last element of list", ll.last().data)

	/* fmt.Println("Clear list")
	ll.clear()
	fmt.Println("Size of list", ll.size()) */

	fmt.Println("Element at 0", ll.getAt(0).data)
	fmt.Println("Element at 2", ll.getAt(2).data)
	fmt.Println("Element at 5", ll.getAt(5))

	fmt.Println("Inserting 7 at index 1")
	ll.insertAt(1,7)
	fmt.Println("Element at 1", ll.getAt(1).data)

	fmt.Println("Inserting 8 at index 0")
	ll.insertAt(0,8)
	fmt.Println("Element at 0", ll.getAt(0).data)

	fmt.Println("Inserting 18 at index 10")
	ll.insertAt(10,18)
	fmt.Println("last element", ll.last().data)

	fmt.Println("Element at 0", ll.getAt(0).data)
	fmt.Println("Element at 1", ll.getAt(1).data)


	fmt.Println("Removing element at index 0")
	ll.removeAt(0)
	fmt.Println("Element at 0", ll.getAt(0).data)

}

func (ll *LinkedList) first() *Node{
	return ll.head
}

func (ll *LinkedList) last() *Node{
	if ll.head == nil {
		return nil
	}

	tempNode := ll.head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}

	return tempNode
}

func (ll *LinkedList) push(data int) {

	node := &Node{data: data}

	if ll.head == nil {
		ll.head = node
		return 
	}

	tempNode := ll.head

	for tempNode.next != nil{
		tempNode = tempNode.next
	}

	tempNode.next = node
}

func (ll *LinkedList) size()(size int){
	size = 0
	tempNode := ll.head
	for tempNode != nil {
		size++
		tempNode = tempNode.next
	}
	return
}

func (ll *LinkedList) clear() {
	ll.head = nil
	return
}

func (ll *LinkedList) insertFirst(data int){
	node := &Node{data: data, next: ll.head}
	ll.head = node
	return
}

func (ll *LinkedList) getAt(index int)(*Node){
	counter := 0
	tempNode := ll.head

	for tempNode != nil {
		if(counter == index){
			return tempNode
		}
		counter++
		tempNode = tempNode.next
	}

	return nil
}

func (ll *LinkedList) insertAt(index int, data int){
	node := &Node{data: data}
	if ll.head == nil {
		ll.head = node
		return
	}

	if(index == 0){
		ll.insertFirst(data)
		return
	}

	previousNode := ll.getAt(index-1) 

	if previousNode == nil {
		previousNode = ll.last()
	}

	tempNode := previousNode.next
	previousNode.next = node
	node.next = tempNode
}

func (ll *LinkedList) removeAt(index int){
	if ll.head == nil{
		return
	}

	if index == 0 {
		ll.head = ll.head.next
	}

	previousNode := ll.getAt(index-1)

	if previousNode != nil {
		previousNode.next = previousNode.next.next
	}
}