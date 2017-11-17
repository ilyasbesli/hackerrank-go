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
	splitted := strings.Split(numberStr, " ")
	s, _ := strconv.Atoi(strings.TrimSpace(splitted[0]))
	t, _ := strconv.Atoi(strings.TrimSpace(splitted[1]))

	numberStr, _ = reader.ReadString('\n')
	splitted = strings.Split(numberStr, " ")
	a, _ := strconv.Atoi(strings.TrimSpace(splitted[0]))
	b, _ := strconv.Atoi(strings.TrimSpace(splitted[1]))

	numberStr, _ = reader.ReadString('\n')
	splitted = strings.Split(numberStr, " ")
	m, _ := strconv.Atoi(strings.TrimSpace(splitted[0]))
	n, _ := strconv.Atoi(strings.TrimSpace(splitted[1]))

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
		aaa, _ := strconv.Atoi(strings.TrimSpace(splitted[i]))

		orangeDistances = append(orangeDistances, aaa)
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
