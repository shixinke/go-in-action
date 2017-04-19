package main

import "fmt"
var score int

func reward(score int) bool{
	if score > 90 {
		return true
	} else {
		return false
	}

}

func main() {
	score = 99
	if score > 90 {
		fmt.Println("优秀")
	} else if score > 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("同学，要努力了哟")
	}
	if isExcellent := reward(score); isExcellent {
		fmt.Println("奖励一个奖牌")
	}
}
