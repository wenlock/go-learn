package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(fmt.Sprintf("%s", runtime.Version()))
}
