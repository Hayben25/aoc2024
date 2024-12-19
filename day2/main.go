package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(window []byte) bool {
	start := 0
	j := 0
	for string(window[j]) != " " {
		j++
	}
	start = j + 1
	prev, err := strconv.Atoi(string(window[0:j]))
	first := true
	if err != nil {
		panic(err)
	}
	increasing := true
	for i := start; i < len(window); i++ {
		if string(window[i]) == " " {
			num, err := strconv.Atoi(string(window[start:i]))
			if err != nil {
				panic(err)
			}
			if math.Abs(float64(num-prev)) > 3 || num == prev {
				return false
			}
			if first {
				increasing = num-prev > 0
				first = false
			} else if increasing != (num-prev > 0) {
				return false
			}
			prev = num
			start = i + 1
		}

	}
	num, err := strconv.Atoi(string(window[start:]))
	if err != nil {
		panic(err)
	}
	if math.Abs(float64(num-prev)) > 3 || num == prev {
		return false
	}
	if first {
		increasing = num-prev > 0
		first = false
	} else if increasing != (num-prev > 0) {
		return false
	}
	return true
}
func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	start := 0
	ct := 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "\n" {
			if check(data[start:i]) {
				fmt.Println("window", string(data[start:i]))
				ct++
			}
			start = i + 1
		}
	}
	fmt.Print(ct)
}
