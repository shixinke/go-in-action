package spider
import "fmt"
func init() {
	fmt.Println("main init")
}

//对外的唯一入口
func Run() {
	//1、下载
	Download("http://www.baidu.com")
	//2、分析
	Analysis()
	//3、存储
	Store("test.txt", "test data")
}
