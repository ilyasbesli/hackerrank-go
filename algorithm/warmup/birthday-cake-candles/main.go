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

	totalNumber, _ := strconv.Atoi(strings.TrimSpace(numberStr))
	numbersArr := strings.Split(numbersStr, " ")
	_ = totalNumber

	blowOut := 0
	max := 0

	for i, str := range numbersArr {
		if i > totalNumber {
			break
		}
		no, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			log.Printf("coudl not parse, %v", err)
		}
		if no > max {
			blowOut = 1
			max = no
		} else if no == max {
			blowOut++
		}
	}

	fmt.Printf("%d\n", blowOut)

}
