package main

import (
	"fmt"
	"os"
	"strconv"
)

func move(x int, y int, arr [][]int, start int) int {
	ct := 0
	fmt.Println(start)
	if y != 0 && arr[y-1][x] == start+1 {
		if arr[y-1][x] == 9 {
			ct++
		} else {
			move(x, y-1, arr, start+1)
		}
	}
	if y != len(arr)-1 && arr[y+1][x] == start+1 {
		if arr[y+1][x] == 9 {
			ct++
		} else {
			move(x, y+1, arr, start+1)
		}
	}
	if x != 0 && arr[y][x-1] == start+1 {
		if arr[y][x-1] == 9 {
			ct++
		} else {
			move(x-1, y, arr, start+1)
		}
	}
	if x != len(arr[y])-1 && arr[y][x+1] == start+1 {
		if arr[y][x+1] == 9 {
			ct++
		} else {
			move(x+1, y, arr, start+1)
		}
	}
	return ct
}

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	arr := make([][]int, 0)
	leng := 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "\n" {
			arr = make([][]int, len(data)/i-1)
			leng = i
			break
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, leng)
	}
	y := 0
	x := 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) != "\n" {
			num, err := strconv.Atoi(string(data[i]))
			if err != nil {
				panic(err)
			}
			arr[y][x] = num
			x++
		} else {
			y++
			x = 0
		}
	}
	ct := 0
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 0 {
				ct += move(j, i, arr, 0)
			}
		}
		fmt.Println()
	}
	fmt.Println(ct)
}
