package main

import "strings"

func main() {
	a := "hahahah"
	b := "heheheh"
	c := strings.Join([]string{a, b}, ",")
	println(c)
}
