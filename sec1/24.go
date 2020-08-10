package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)

	*p = 4
	fmt.Println(*p)

}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var d int
	return &d
}

func delta(old, new int) int {
	return new - old
}
