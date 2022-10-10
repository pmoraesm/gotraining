package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Input a string:")

	// Using bufio instead of fmt to correctly capture input with spaces.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	myString := scanner.Text()
	s := strings.ToLower(myString)

	switch {
	case strings.HasPrefix(s, "i") == false:
		fmt.Printf("Not Found!\n")
	case strings.HasSuffix(s, "n") == false:
		fmt.Printf("Not Found!\n")
	case strings.ContainsAny(s, "a") == false:
		fmt.Printf("Not Found!\n")
	default:
		fmt.Printf("Found!\n")
	}
}
