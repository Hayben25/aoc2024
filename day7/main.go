package main

import (
	"fmt"
	"os"
	"strconv"
)

type block struct {
	nums string
}

func check(data []byte) int {
	i := 0
	for string(data[i]) != ":" {
		i++
	}
	num, err := strconv.Atoi(string(data[0:i]))
	if err != nil {
		panic(err)
	}
	i += 2
	num_ct := 1
	start := i
	for i < len(data) {
		if data[i] == 32 {
			num_ct++
		}
		i++
	}
	nums := make([]int, num_ct)
	num_ct = 0
	i = start
	for i < len(data) {
		if data[i] == 32 {
			num_temp, err := strconv.Atoi(string(data[start:i]))
			if err != nil {
				panic(err)
			}
			nums[num_ct] = num_temp
			num_ct++
			start = i + 1
		}
		i++
	}
	num_temp, err := strconv.Atoi(string(data[start:]))
	if err != nil {
		panic(err)
	}
	nums[len(nums)-1] = num_temp
	sum := 0
	if death(num, nums) {
		for i = 0; i < len(nums); i++ {
			sum += nums[i]
		}
	}
	return sum
}
func death(num int, nums []int) bool {
	for i := 2; i < 1<<len(nums)-1; i++ {
		fmt.Println(i, num)
	}
	return false
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
			ct += check(data[start:i])
			start = i + 1
		}
	}
	fmt.Println(ct)
}
