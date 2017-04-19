package main

import "fmt"

type Programming struct {
	desc string
	version string
	ref   int
}

type Web interface {
	request() map[string] string
	response() string
}

type PHP struct {
	//有点类似于面向对象中的继承，但是却又不一样（有一点像PHP中的traits）
	Programming
	hashTables map[string] string
}

func (w *Programming) getRef() int{
	return w.ref
}

func (p *PHP) getRef() int {
	return p.ref+1
}

func (p *PHP) request() map[string]string {
	return map[string]string{"charset":"utf-8", "url":"/"}
}

func (p *PHP) response() string {
	return "server running"
}

func (p *PHP) echo(str string) {
	fmt.Println(p.desc+" echo:"+str)
}

func main() {
	phpweb := PHP{Programming{"PHP", "7.1", 1}, map[string]string{"engine":"zend"}}
	fmt.Println(phpweb.desc)
	fmt.Println(phpweb.version)
	fmt.Println(phpweb.hashTables["engine"])
	phpweb.echo("web programming")

	//方法重写
	fmt.Println(phpweb.ref)
	fmt.Println(phpweb.getRef())

	//phpweb实现了web这个接口所有的方法，那么就说明phpweb实现了web这个接口
	fmt.Println(phpweb.request()["url"])
	fmt.Println(phpweb.response())
}
