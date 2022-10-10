package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myNo float64
	fmt.Printf("Input a float:")
	fmt.Scan(&myNo)
	sInt := strconv.FormatInt(int64(myNo), 10)
	fmt.Printf("Truncated version: %v\n", sInt)
}
