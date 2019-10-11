package main

import "fmt"

func p() string {
	a := "测试"
	fmt.Println(a)
	return a
}

func max(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(max(1, 2))
}
