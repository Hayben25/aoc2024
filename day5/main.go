package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lis struct {
	pt1 map[int]string
}

func (pt1 *lis) check(line string) int {
	line_list := make([]int, (len(line)+1)/3)
	for i := 0; i < len(line); i++ {
		if string(line[i]) == "," {
			num, err := strconv.Atoi(string(line[i-2 : i]))
			if err != nil {
				panic(err)
			}
			line_list[i/3] = num
		}
	}
	num, err := strconv.Atoi(string(line[len(line)-2:]))
	if err != nil {
		panic(err)
	}
	line_list[len(line_list)-1] = num
	for i := 0; i < len(line_list); i++ {
		for j := i + 1; j < len(line_list); j++ {
			str := strconv.Itoa(line_list[i])
			if strings.Contains(pt1.pt1[line_list[j]], str) {
				return 0
			}
		}
	}
	return line_list[len(line_list)/2]
}
func main() {
	p := 0
	pt1 := 0
	data, err := os.ReadFile("actual.txt")
	if err != nil {
		panic(err)
	}
	for p < len(data) {
		if string(data[p]) == "\n" {
			if string(data[p-3]) != "|" {
				pt1 = p - 1
				break
			}
		}
		p++
	}
	pt1_str := string(data[:pt1])
	pt1_map := map[int]string{}
	for i := 0; i < len(pt1_str); i++ {
		if string(pt1_str[i]) == "\n" {
			num1, err := strconv.Atoi(string(pt1_str[i-5 : i-3]))
			if err != nil {
				panic(err)
			}

			pt1_map[num1] += string(pt1_str[i-2:i]) + " "
		}
	}
	pt2_str := string(data[pt1+2 : len(data)-1])
	pt1_list := lis{pt1: pt1_map}
	start := 0
	ct := 0
	for i := 0; i < len(pt2_str); i++ {
		if string(pt2_str[i]) == "\n" {
			ct += pt1_list.check(string(pt2_str[start:i]))
			start = i + 1
		}
	}
	fmt.Println(ct)
}
