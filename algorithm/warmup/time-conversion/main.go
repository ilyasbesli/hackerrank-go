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
	s, _ := reader.ReadString('\n')
	timeStr := []rune(s)

	// 07:05:45PM
	hourStr := string(timeStr[0:2])
	minStr := string(timeStr[3:5])
	secStr := string(timeStr[6:8])
	p := strings.TrimRight(string(timeStr[len(s)-2:]), "\n")

	hour, _ := strconv.Atoi(strings.TrimSpace(string(hourStr)))

	if p == "PM" {
		if hour == 12 {
			hour = 12
		} else {
			hour = hour + 12
		}
	} else if p == "AM" {
		if hour == 12 {
			hour = 0
		}
	}

	if hour > 9 {
		fmt.Printf("%d:%s:%s\n", hour, minStr, secStr)
	} else {
		fmt.Printf("0%d:%s:%s\n", hour, minStr, secStr)
	}

}
