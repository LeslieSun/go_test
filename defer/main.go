package main

import "fmt"

func f1() (r int) {
	r = 1
	defer func() {
		r++
		fmt.Println("r value =", r)
	}()
	r = 2
	return r
}
func f2() {
	i := 1
	defer fmt.Println("Derferred print:", i)
	i++
	fmt.Println("Normal print:", i)
}
func f3() int {
	i := 1
	defer func() {
		i++
		fmt.Println("Derferred print:", i)
	}()
	return i
}
func main() {
	f1()
	//f2()
	//i := f3()
	//fmt.Println("outer print:", i)
}
