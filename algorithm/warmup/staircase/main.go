package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	numberStr, _ := reader.ReadString('\n')
	totalNumber, _ := strconv.Atoi(strings.TrimSpace(numberStr))

	for i := 0; i < totalNumber; i++ {
		for j := 0; j < totalNumber-i-1; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i+1; j++ {
			fmt.Print("#")
		}

		fmt.Println()
	}
}
