package spider

import "fmt"

func Store(filename string, data string) {
	fmt.Printf("存储文件为:%s,存储内容为:%s\n", filename, data)
}

func init() {
	fmt.Println("store init")
}
