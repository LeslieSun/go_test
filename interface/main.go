package main

type IPeople interface {
	hello()
}
type People struct {
	str string
}

func (p *People) hello() {
}
func test(i IPeople) {
}
func main() {
	val := People{"123"}
	test(val)
	test(&val)
}
