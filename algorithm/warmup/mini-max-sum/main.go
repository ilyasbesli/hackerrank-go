package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(arr []int64) int64 {
	var su int64
	su = 0
	for i := 0; i < len(arr); i++ {
		su = su + arr[i]
	}

	return su
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	numbersStr, _ := reader.ReadString('\n')
	numbersArr := strings.Split(numbersStr, " ")

	var numbers []int64

	for _, str := range numbersArr {
		no, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 64)

		numbers = append(numbers, no)
	}

	var min, max int64

	min = math.MaxInt64
	max = math.MinInt64

	for i := 0; i < len(numbers); i++ {
		b := make([]int64, len(numbers))
		copy(b, numbers)
		a := append(b[:i], b[i+1:]...)
		sumOfArray := sum(a)

		if min > sumOfArray {
			min = sumOfArray
		}

		if max < sumOfArray {
			max = sumOfArray
		}
	}

	fmt.Printf("%d %d\n", min, max)
}
