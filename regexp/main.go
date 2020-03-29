package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `begin_time:2020-01-13 02:36:18
	end_time:2020-01-13 02:38:56
	diff:158s`
	fmt.Println(str)
	r, _ := regexp.Compile("diff:[0-9]*s")
	res := r.FindStringSubmatch(str)
	r1 := regexp.MustCompile(`[0-9]{1,}`)
	fmt.Println(res[0])
	params := r1.FindStringSubmatch(res[0])
	for i, param := range params {
		fmt.Println(i, ",", param)
	}
}
