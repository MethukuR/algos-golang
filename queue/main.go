package main

import (
	"fmt"
	"math"
)
type Queue struct {
	value []int
	size int
}

func main(){
	queue := &Queue{}

	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	queue.printQueue()

	fmt.Println()

	fmt.Println(queue.Remove())

	queue.printQueue()

	fmt.Println()

	fmt.Println(queue.Top())

	queue.printQueue()

	fmt.Println(queue.Remove())

	fmt.Println(queue.Remove())

	fmt.Println(queue.Remove())

	fmt.Println(queue.Remove())

}

func (q *Queue) Add(data int){
	q.value = append(q.value,data)
}

func (q *Queue) Remove()(top int){
	if(len(q.value) == 0 ){
		return math.MinInt64
	}
	
	top = q.value[0]
	q.value = q.value[1:]
	return 
}

func (q *Queue) Top()(top int){
	if(len(q.value) == 0 ){
		return math.MinInt64
	}

	top = q.value[0]
	return 
}

func (q *Queue) printQueue(){
	for i := 0 ; i < len(q.value) ; i++ {
		fmt.Print(q.value[i])
	}

	fmt.Println()
}