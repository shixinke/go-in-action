package main
import "fmt"

type Person struct {
	name string
	sex string
	age int
}

func main() {
	//第一种
	var p Person
	p.name = "shixinke"
	p.sex = "M"
	p.age = 29
	//第二种
	p2 := new(Person)
	p2.name = "golang"
	//第三种
	p3 := Person {name:"Hanmei", sex:"F", age :28}
	fmt.Printf("姓名:%s, 性别:%s, 年龄:%d\n", p.name, p.sex, p.age)
	fmt.Printf("姓名:%s, 性别:%s, 年龄:%d\n", p2.name, p2.sex, p2.age)
	fmt.Printf("姓名:%s, 性别:%s, 年龄:%d\n", p3.name, p3.sex, p3.age)
}
