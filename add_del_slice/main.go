package main

import "fmt"

func input(stack []int) {
	stack = append(stack, 4)
}
func main() {
	sli := []int{1, 2, 3}
	fmt.Printf("in main sli:%v", sli)
	input(sli)
	fmt.Printf("in main sli:%v", sli)
}
