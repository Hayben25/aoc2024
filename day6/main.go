package main

import (
	"fmt"
	"os"
)

type gaurd struct {
	x         int
	y         int
	direction string
}

func (g *gaurd) move(arr [][]string) {
	arr[g.y][g.x] = "X"
	if g.direction == "up" {
		if arr[g.y-1][g.x] == "#" {
			g.direction = "right"
		} else {
			g.y--
		}
	}
	if g.direction == "right" {
		if arr[g.y][g.x+1] == "#" {
			g.direction = "down"
		} else {
			g.x++
		}
	}
	if g.direction == "down" {
		if arr[g.y+1][g.x] == "#" {
			g.direction = "left"
		} else {
			g.y++
		}
	}
	if g.direction == "left" {
		if arr[g.y][g.x-1] == "#" {
			g.direction = "up"
		} else {
			g.x--
		}
	}
}
func main() {
	data, err := os.ReadFile("actual.txt")
	if err != nil {
		panic(err)
	}
	line_ct := 1
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "\n" {
			line_ct++
		}
	}
	arr := make([][]string, line_ct-1)
	line_ct -= 2
	for line_ct > -1 {
		arr[line_ct] = make([]string, 130)
		line_ct--
	}
	k := 0
	g := gaurd{direction: "up"}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if string(data[k]) == "\n" {
				k++
			}
			if string(data[k]) == "^" {
				g.x = j
				g.y = i
			}
			arr[i][j] = string(data[k])
			k++
		}
	}
	fmt.Print(arr)
	for g.x != 0 && g.x != len(arr[0])-1 && g.y != 0 && g.y != len(arr)-1 {
		g.move(arr)
	}
	ct := 0
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == "X" {
				ct++
			}
		}
	}
	fmt.Println(ct + 1)
}
