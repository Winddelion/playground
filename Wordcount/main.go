package main

import "fmt"
import "os"
import "bufio"
import "strings"

func CountWords(arr string) int16 {
	dict := strings.Split(arr, " ")
	if len(dict) == 1 && dict[0] == "" {
		return 0
	}
	return int16(len(dict))
}

func main() {

	fmt.Println("Type any amount of words you would like.")
	fmt.Printf("> ")

	UserInput := bufio.NewScanner(os.Stdin)
	UserInput.Scan()

	fmt.Printf("Your number of words is: %v\n", CountWords(UserInput.Text()))
}
