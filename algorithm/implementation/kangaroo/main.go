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
	slittedString := strings.Split(numberStr, " ")

	if len(slittedString) != 4 {
		log.Fatalf("could not parse, you need at least 4 parameter, but we got %v", slittedString)
	}

	start1 := parseOrExit(slittedString[0])
	step1 := parseOrExit(slittedString[1])

	start2 := parseOrExit(slittedString[2])
	step2 := parseOrExit(slittedString[3])

	// if it is impossible to catch.
	if start2 > start1 && step2 > step1 {
		fmt.Println("NO")
	} else {
		for i := 0; i < 10000; i++ {
			n1 := start1 + step1*i
			n2 := start2 + step2*i
			if n1 == n2 {
				fmt.Println("YES")
				os.Exit(0)
			}
		}

		fmt.Println("NO")
	}
}

func parseOrExit(str string) int {
	no, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Printf("could not parse, %v", err)
		os.Exit(1)
	}

	return no
}
