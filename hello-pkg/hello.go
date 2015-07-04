package hello

import (
	"fmt"
	"strings"
)

//Yell public function
func Yell(s string) {
	fmt.Println(fmt.Sprintf("HELLO %s", strings.ToUpper(s)))
}
