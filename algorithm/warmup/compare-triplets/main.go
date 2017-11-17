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

	aliceNumbersStr, _ := reader.ReadString('\n')
	bobNumbersStr, _ := reader.ReadString('\n')

	aliceNumber := strings.Split(aliceNumbersStr, " ")
	bobArrNumber := strings.Split(bobNumbersStr, " ")

	aliceScore := 0
	bobScore := 0

	for ind := range aliceNumber {
		firstNo, _ := strconv.Atoi(strings.TrimSpace(aliceNumber[ind]))
		secondNo, _ := strconv.Atoi(strings.TrimSpace(bobArrNumber[ind]))

		diff := firstNo - secondNo

		if diff > 0 {
			aliceScore++
		} else if diff < 0 {
			bobScore++
		}

	}

	fmt.Printf("%d %d \n", aliceScore, bobScore)
}
