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

	var numbers []int

	for i := 0; i < totalNumber; i++ {
		str, _ := reader.ReadString('\n')
		no, _ := strconv.Atoi(strings.TrimSpace(str))

		numbers = append(numbers, no)
	}

	for i := 0; i < totalNumber; i++ {
		no := numbers[i]
		t := no % 5

		realNumber := no
		if t >= 3 && no >= 38 {
			realNumber = no + 5 - t
		}

		fmt.Printf("%d\n", realNumber)
	}

}
