package main

func main() {

}
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
