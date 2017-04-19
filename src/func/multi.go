package main

import "fmt"

var students [20]int
func result(arr [20]int) (int, int) {
	for i :=0; i< len(arr); i++ {
		if arr[i] == 90 {
			return i,arr[i]
		}
	}
	return 0,0
}

func main() {
	for i :=1 ;i < 20 ; i++ {
		students[i] = i*5
	}
	var index, score = result(students)
	fmt.Printf("获得优秀的同学的学号为%d,成绩为%d\n", index, score)
}
