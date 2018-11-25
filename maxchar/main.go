// --- Directions
// Given a string, return the character that is most
// commonly used in the string.
// --- Examples
// maxChar("abcccccccd") === "c"
// maxChar("apple 1231111") === "1"

package main

import (
	"fmt"
	"strings"
)

func main(){
	inputStr := "apple 1231111"
	maxChar := getMaxChar(inputStr)
	fmt.Println(maxChar)
}

func getMaxChar(inputStr string)(maxChar string){
	charMap := make(map[string]int)
	strArr := strings.Split(inputStr, "")

	for _, char := range strArr {
		if(charMap[char] == 0){
			charMap[char] = 1
		}else{
			charMap[char] += 1
		}
	}

	maxCharValue := 0
	for char , value := range charMap { 
		if(value > maxCharValue){
			maxChar = char
			maxCharValue = value
		}
	}
	return
}