package main

import (
	"bufio"
	"fmt"
	"os"
)

func reduce(str []rune) []rune {
	if len(str) == 0 {
		return str
	}
	first := str[0]
	count := 0
	for i, v := range str {
		if v == first {
			count++
		} else {

			return append(reduce(str[i:]), first)
		}

		if count == 3 {
			newRune := str[3:]
		}

	}

	return str
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')

	r := []rune(str[0:])

	reducedStr := reduce(r)

	if len(reducedStr) == 0 {
		reducedStr = "Empty String"
	}

	fmt.Printf("%s\n", reducedStr)
}
