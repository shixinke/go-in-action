package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(10)
	switch num {
	case 2:
		fmt.Println("得到一张价值为10元的优惠券")
		//break
	case 6:
		fmt.Println("得到一台笔记本电脑")
		//break
	case 8:
		fmt.Println("得到现金8000元")
		//break
	default:
		fmt.Printf("您的数字是%d,您与幸运女神擦肩而过\n", num)
	}
}
