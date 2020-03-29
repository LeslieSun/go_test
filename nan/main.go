package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fee := ""
	f, err := strconv.ParseFloat(fee, 32)
	fmt.Println(f, math.IsNaN(f), err)
	i, err := strconv.ParseInt(fee, 10, 32)
	fmt.Println(i, err)

}
