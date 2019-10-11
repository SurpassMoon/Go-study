package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	p := 1
	var a int
	&a = &p

	fmt.Println(a)

}
