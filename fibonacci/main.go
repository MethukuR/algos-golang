// --- Directions
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3

package main

import (
	"fmt"
)

func main(){
	input := 5
	fmt.Println(fib(input, 1))
}

func fib(input int, result int)(int){

	if input <= 1 {
		return result
	}

	return fib(input-1 , result*input)
}
