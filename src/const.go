package main
import "fmt"

const SUCCESS = true                      //布尔型常量
const FAILED = false
const MAX_LIMIT = 20                     //整型常量
const RATE = 0.86                        //浮点型常量
const MESSAGE = "success"                //字符串常量
const CH = 'w'+1                         //tune常量
const COM = complex(0, 3.75)        //复数常量


func main() {
	fmt.Println(SUCCESS)
	fmt.Println(FAILED)
	fmt.Println(MAX_LIMIT)
	fmt.Println(RATE)
	fmt.Println(MESSAGE)
	fmt.Println(CH)
	fmt.Println(COM)
}
