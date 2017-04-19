package main
import "fmt"

//布尔类型
var finished bool = false
//数字型
var level int = 3
var kpi   float32 = 3.3
//字符串型
var desc string = "description"
//数组(32个类型为byte的数组)
var arr [32]byte
//切片(可以认为是动态数组，即元素个数不固定的数组)
var multi []string
//结构体
var student struct {
	stuNo int
	stuName string
}
//指针
var ptr *int
//函数类型
var say func(msg string) string
//接口类型
var plugin interface {}
//映射类型
var obj map[string]string
//管道类型
var ch chan int
//自定义类型
type Age int
var age Age

func main() {
	age = 29
	fmt.Println(finished)
	fmt.Println(level)
	fmt.Println(kpi)
	fmt.Println(desc)

	fmt.Println(ch)
	fmt.Println(age)
}
