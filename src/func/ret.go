package main

import "fmt"

func sumAndMul(num1 int, num2 int) (sum int, mul int) {
	sum = num1 + num2
	mul = num1 * num2
	return
}
func main() {
	a, b := 10, 20
	sum, mul := sumAndMul(a, b)
	fmt.Printf("%d + %d = %d, %d * %d = %d \n", a, b, sum, a, b, mul)
}
