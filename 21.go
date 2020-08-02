package main

import "fmt"

func main() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	p := f()

	fmt.Println(p, *p)

	fmt.Println(f() == f()) // "false"

}

func f() *int {
	v := 1
	return &v
}
