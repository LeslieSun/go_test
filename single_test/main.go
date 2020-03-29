package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"time"
	"unsafe"
)

func sum(vars ...int) int {
	total := 0
	for _, v := range vars {
		total += v
	}
	return total
}
func ff(out io.Writer) {
	if out != nil {
		out.Write([]byte("done\n"))
	}
}

type empty_inter interface {
}
type nono_empty_inter interface {
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	s := []int{2, 3, 4, 5}
	fmt.Println(sum(s...))

	const day = time.Hour * 24
	fmt.Println("Type:", reflect.TypeOf(day).String())

	var i empty_inter
	fmt.Println("size of empty interface:", unsafe.Sizeof(i))
	var j nono_empty_inter
	fmt.Println("size of none empty interface:", unsafe.Sizeof(j))

	var buf bytes.Buffer
	ff(buf)
}
