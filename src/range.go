package main

import "fmt"
var scores [20]int

func main() {
	for i := 0; i<20; i++ {
		scores[i] = 20*i
	}
	for _, num := range scores {
		fmt.Println(num)
	}
}
