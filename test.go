package main

import "fmt"

type test1 struct {
	name int
	age  int
}

type test2 struct {
	name int
	age  int
}

func main() {
	a := test1{}
	b := test2{}
	fmt.Println(a == b)
}
