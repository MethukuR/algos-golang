package main

type Node struct{
	data int
	neighbours []*Node
	visited bool
	startTime int
	endTime int
}