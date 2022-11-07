package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var input string

	sli := make([]int, 0, 3)

	for input != "X" {
		fmt.Printf("Enter a number:\n")
		fmt.Scan(&input)

		x, err := strconv.Atoi(input)
		if errors.Is(err, strconv.ErrSyntax) {
			continue
		}

		sli = append(sli, x)
		sort.Sort(sort.IntSlice(sli))
		fmt.Printf("%v\n", sli)
	}
}
