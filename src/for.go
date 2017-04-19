package main

import "fmt"

func main() {
	//第一种，类似于while
	var num int = 0
	var sum int = 0
	for num < 10 {
		sum += num
		num ++
	}
	fmt.Printf("总和为:%d\n", sum)
	//第二种，和C的for类似
	var product int = 1
	for i := 1; i< 10; i++ {
		product *= i
	}
	fmt.Printf("乘积为:%d\n", product)
	//第三种，死循环
	for {
		fmt.Println("dead")
	}
}
