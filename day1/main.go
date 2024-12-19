package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}
type BST struct {
	root *Node
}

func (bst *BST) Insert(val int) {
	bst.InsertRec(bst.root, val)
}
func (bst *BST) InsertRec(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val <= node.data {
		node.left = bst.InsertRec(node.left, val)
	}
	if val > node.data {
		node.right = bst.InsertRec(node.right, val)
	}
	return node
}
func (bst *BST) ToString(node *Node) string {
	if node == nil {
		return ""
	}
	left := bst.ToString(node.left)
	right := bst.ToString(node.right)
	return left + strconv.Itoa(node.data) + right
}
func (bst *BST) ToList() [1000]int {
	list := [1000]int{}
	bst_string := bst.ToString(bst.root)
	i := 0
	j := 0
	for i < 1000 {
		num, err := strconv.Atoi(bst_string[j : j+5])
		if err != nil {
			panic(err)
		}
		list[i] = num
		i++
		j += 5
	}
	return list
}
func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	i := 14
	list1 := BST{}
	list2 := BST{}
	for i <= len(data) {
		num, err := strconv.Atoi(string(data[i-14 : i-9]))
		if err != nil {
			panic(err)
		}
		list1.Insert(num)
		num, err = strconv.Atoi(string(data[i-6 : i-1]))
		if err != nil {
			panic(err)
		}
		list2.Insert(num)
		i += 14
	}
	list1_sorted := [1000]int{}
	list2_sorted := [1000]int{}

	list1_sorted = list1.ToList()
	list2_sorted = list2.ToList()
	sum := int64(0)
	for i = 0; i < 1000; i++ {
		sum += int64(max(list1_sorted[i], list2_sorted[i]) - min(list1_sorted[i], list2_sorted[i]))
	}
	fmt.Print(sum)
}
