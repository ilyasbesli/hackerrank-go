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
	numbersStr, _ := reader.ReadString('\n')

	totalNumber := parseOrExit(numberStr)
	numbersArr := strings.Split(numbersStr, " ")

	possitiveNumberCount := 0
	negativeNumberCount := 0
	zeroNumberCount := 0

	for _, str := range numbersArr {
		no := parseOrExit(str)

		if no < 0 {
			negativeNumberCount++
		} else if no > 0 {
			possitiveNumberCount++
		} else {
			zeroNumberCount++
		}
	}

	fmt.Printf("%v\n", float32(possitiveNumberCount)/float32(totalNumber))
	fmt.Printf("%v\n", float32(negativeNumberCount)/float32(totalNumber))
	fmt.Printf("%v\n", float32(zeroNumberCount)/float32(totalNumber))

}

func parseOrExit(str string) int {
	no, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		log.Printf("could not parse, %v", err)
		os.Exit(1)
	}

	return no
}
