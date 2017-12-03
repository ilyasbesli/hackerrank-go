package main

import (
	"bufio"
	"fmt"
	"log"
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
		firstNo := parseOrExit(aliceNumber[ind])
		secondNo := parseOrExit(bobArrNumber[ind])

		diff := firstNo - secondNo

		if diff > 0 {
			aliceScore++
		} else if diff < 0 {
			bobScore++
		}

	}

	fmt.Printf("%d %d \n", aliceScore, bobScore)
}

func parseOrExit(str string) int {
	no, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Printf("could not parse, %v", err)
		os.Exit(1)
	}

	return no
}
