package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(data []byte, i int) (int, int) {
	if string(data[i:i+4]) != "mul(" {
		return 0, i
	}
	i += 4
	ind_comma, ind_para := 0, 0
	fmt.Println(string(data[i:]))
	for k := i; k <= i+8; k++ {
		if string(data[k]) == "," {
			ind_comma = k
		} else if string(data[k]) == ")" {
			ind_para = k
			break
		}
	}
	if ind_comma-i == 0 || ind_para-i == 0 || ind_para-i > 8 || ind_comma-i > 4 || ind_comma-i < 1 || ind_para-i < 3 {
		return 0, i
	}
	num, err := strconv.Atoi(string(data[i:ind_comma]))
	if err != nil {
		panic(err)
	}
	num1, err := strconv.Atoi(string(data[ind_comma+1 : ind_para]))
	if err != nil {
		panic(err)
	}
	fmt.Print(num)
	fmt.Print(" ")
	fmt.Println(num1)
	return num * num1, i + ind_para
}
func main() {
	data, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for i := 0; i < len(data); i++ {
		mul, loc := check(data, i)
		total += mul
		i = loc
	}
	fmt.Print(total)
	fmt.Println(" Correct Answer:: 180233229")
}
