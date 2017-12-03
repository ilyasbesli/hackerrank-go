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

	numberStr, _ := reader.ReadString('\n')
	splitted := strings.Split(numberStr, " ")
	s := parseOrExit(splitted[0])
	t := parseOrExit(splitted[1])

	numberStr, _ = reader.ReadString('\n')
	splitted = strings.Split(numberStr, " ")
	a := parseOrExit(splitted[0])
	b := parseOrExit(splitted[1])

	numberStr, _ = reader.ReadString('\n')
	splitted = strings.Split(numberStr, " ")
	m := parseOrExit(splitted[0])
	n := parseOrExit(splitted[1])

	numberStr, _ = reader.ReadString('\n')
	splitted = strings.Split(numberStr, " ")
	var appleDistances []int

	for i := 0; i < m; i++ {
		aaa, _ := strconv.Atoi(strings.TrimSpace(splitted[i]))

		appleDistances = append(appleDistances, aaa)
	}

	numberStr, _ = reader.ReadString('\n')
	splitted = strings.Split(numberStr, " ")
	var orangeDistances []int

	for i := 0; i < n; i++ {
		dist := parseOrExit(splitted[i])

		orangeDistances = append(orangeDistances, dist)
	}

	numberOfApple := 0
	numberOfOrange := 0

	for i := 0; i < len(appleDistances); i++ {
		x := a + appleDistances[i]

		if x >= s && x <= t {
			numberOfApple++
		}
	}

	for i := 0; i < len(orangeDistances); i++ {
		x := b + orangeDistances[i]

		if x >= s && x <= t {
			numberOfOrange++
		}
	}

	fmt.Printf("%d\n%d\n", numberOfApple, numberOfOrange)
}

func parseOrExit(str string) int {
	no, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Printf("could not parse, %v", err)
		os.Exit(1)
	}

	return no
}
