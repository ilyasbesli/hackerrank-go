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
	numbersStr, _ := reader.ReadString('\n')

	totalNumber, _ := strconv.Atoi(strings.TrimSpace(numberStr))
	numbersArr := strings.Split(numbersStr, " ")

	possitiveNumberCount := 0
	negativeNumberCount := 0
	zeroNumberCount := 0

	for _, str := range numbersArr {
		no, _ := strconv.Atoi(strings.TrimSpace(str))

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
