// switch多路选择

switch coinflip() {
case "heads":
    heads++
case "tails":
    tails++
default:
    fmt.Println("landed on edge!")
}


// 无tag switch(tagless switch)
func Signum(x int) int {
    switch {
    case x > 0:
        return +1
    default:
        return 0
    case x < 0:
        return -1
    }
}

type Point struct {
    X, Y int
}
var p Point
