// --- Directions
// Given a string, return a new string with the reversed
// order of characters
// --- Examples
//   reverse('apple') === 'leppa'
//   reverse('hello') === 'olleh'
//   reverse('Greetings!') === '!sgniteerG'

package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "apple"
	fmt.Println(reverseString(str))
}

func reverseString(str string)(reversedString string) {
	strArr := strings.Split(str, "")
	i := 0
	j := len(strArr) -1

	for i<j {
		strArr[i],strArr[j] = strArr[j], strArr[i]
		i++
		j--
	}
	reversedString = strings.Join(strArr, "")
	return 
}