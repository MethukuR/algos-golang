// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers.
// --- Examples
//   reverseInt(15) === 51
//   reverseInt(981) === 189
//   reverseInt(500) === 5
//   reverseInt(-15) === -51
//   reverseInt(-90) === -9

package main

import (
	"fmt"
	"strings"
	"math"
	"strconv"
)

func main(){
	inpNum := -90
	fmt.Println(revereInt(inpNum))
}

func revereInt(inpNum int)(reveredInt int){
	absValue := math.Abs(float64(inpNum))
	intToString := strconv.Itoa(int(absValue))
	reversedString := reverseString(intToString)
	reveredInt, err := strconv.Atoi(reversedString)

	if err == nil {
		if inpNum <= 0 {
			reveredInt *= -1
		}
	} 
	return
}

func reverseString(str string)(reversedString string) {
	strArr := strings.Split(str, "")
	i := 0
	j := len(strArr) -1

	for i<j {
		strArr[i],strArr[j] = strArr[j], strArr[i]
		i = i+1
		j = j-1
	}
	reversedString = strings.Join(strArr, "")
	return 
}