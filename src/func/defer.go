package main

import "fmt"

func steps() {
	defer fmt.Println("我是第一个来的")
	fmt.Println("我是第二个来的")
	fmt.Println("我是第三个来的")
}

func main() {
	steps()
}
