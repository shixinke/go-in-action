package main

import "fmt"

func add(nums... int) int {
	var total int = 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Printf("1~5的和:%d\n", add(1, 2, 3, 4, 5))
	fmt.Printf("10~13的和:%d\n", add(10, 11, 12, 13))
}
