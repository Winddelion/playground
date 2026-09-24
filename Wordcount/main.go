package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {

	fmt.Println("Type any amount of words you would like.")
	fmt.Printf("> ")

	UserInput := bufio.NewScanner(os.Stdin)
	UserInput.Scan()

	WordList := strings.Split(UserInput.Text(), " ")

	fmt.Printf("Your number of words is: %v\n", len(WordList))
}
