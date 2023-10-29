package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	s, separator := "", ""
	for _, v := range os.Args[1:] {
		s += separator + v
		separator = " "
	}
	fmt.Println(s)
	endTime := time.Now()

	fmt.Println(endTime.Sub(startTime))

	startTime = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	endTime = time.Now()

	fmt.Println(endTime.Sub(startTime))
}
