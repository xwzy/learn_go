package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Arg[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.FPrintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
