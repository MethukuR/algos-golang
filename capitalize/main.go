// --- Directions
// Write a function that accepts a string.  The function should
// capitalize the first letter of each word in the string then
// return the capitalized string.
// --- Examples
//   capitalize('a short sentence') --> 'A Short Sentence'
//   capitalize('a lazy fox') --> 'A Lazy Fox'
//   capitalize('look, it is working!') --> 'Look, It Is Working!'

package main
import (
	"fmt"
	"strings"
)

func main(){
	input := "look, it is working!"
	capitalizedWord := capitalizeWord(input);
	fmt.Println(capitalizedWord)
}

func capitalizeWord(inputStr string)(string){
	words := strings.Split(inputStr, " ")
	capitalizedWords := []string{}
	for _ , word := range words {
		capitalizedWords = append(capitalizedWords, strings.ToUpper(word[0:1])+word[1:])
	}
	return strings.Join(capitalizedWords, " ")
}