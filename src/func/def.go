package main

import "fmt"

func sum(a int, b int) int {
	return a+b
}

func main() {
	num1, num2 := 20, 30
	fmt.Println(sum(num1, num2))
}
