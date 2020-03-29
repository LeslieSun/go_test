package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func printScannerType() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("scanner tpye:%T\n", scanner)
}
func main() {
	printScannerType()
	str := "hello,你好"
	fmt.Println("rune:", utf8.RuneCountInString(str))
}
