package main

import (
	"fmt"
	"strconv" //for conversions to and from string
)

type Human struct {
	name  string
	age   int
	phone string
}

//Returns a nice string representing a Human
//With this method, Human implements fmt.Stringer
func (h Human) String() string {
	//We called strconv.Itoa on h.age to make a string out of it.
	//Also, thank you, UNICODE!
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
