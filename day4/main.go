package main

import (
	"fmt"
	"os"
)

type position struct {
	x_val    int
	y_val    int
	line_len int
	col_size int
}

func (pos *position) check(data []byte, i int) int {
	ct := 0
	if pos.y_val >= 3 {
		if string(data[i-pos.line_len-1]) == "M" && string(data[i-2*pos.line_len-2]) == "A" && string(data[i-3*pos.line_len-3]) == "S" {
			ct++
		}
		if pos.x_val >= 3 && string(data[i-pos.line_len-2]) == "M" && string(data[i-2*pos.line_len-4]) == "A" && string(data[i-3*pos.line_len-6]) == "S" {
			ct++
		}
		if i <= len(data)-5 {
			if string(data[i-pos.line_len]) == "M" && string(data[i-2*pos.line_len]) == "A" && string(data[i-3*pos.line_len]) == "S" {
				ct++
			}
		}
	}
	if pos.col_size-pos.y_val >= 4 {
		if string(data[i+pos.line_len+1]) == "M" && string(data[i+2*pos.line_len+2]) == "A" && string(data[i+3*pos.line_len+3]) == "S" {
			ct++
		}

		if string(data[i+pos.line_len+2]) == "M" && string(data[i+2*pos.line_len+4]) == "A" && string(data[i+3*pos.line_len+6]) == "S" {
			ct++
		}
		if string(data[i+pos.line_len]) == "M" && string(data[i+2*pos.line_len]) == "A" && string(data[i+3*pos.line_len]) == "S" {
			ct++
		}
	}
	if pos.x_val >= 3 {
		if string(data[i-3:i]) == "SAM" {
			ct++
		}
	}
	if string(data[i+1:i+4]) == "MAS" {
		ct++
	}
	return ct
}
func main() {
	data, err := os.ReadFile("actual.txt")
	if err != nil {
		panic(err)
	}
	ct := 0
	i := 0
	pos := position{line_len: 140, col_size: 140}
	for i < len(data) {
		if string(data[i]) == "X" {
			ct += pos.check(data, i)
		}
		i++
		pos.x_val++
		if pos.x_val > pos.line_len {
			pos.x_val = 0
			pos.y_val++
		}
	}
	fmt.Println(ct)
}
