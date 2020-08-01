package main


import "fmt"
import "os"


func main() {

	fmt.Println(os.Args[0])
	
	for idx, value := range os.Args {
		fmt.Println(idx, value)
	}

}