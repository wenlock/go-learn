package main


import (
	"fmt"
)

func plus(a int, b int) int {
	return a+b
}

func main() {

	ret := plus(5,5)
	fmt.Println("sum:",ret)
}