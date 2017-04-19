package spider

import "fmt"

func init() {
	fmt.Println("download init")
}

func Download(url string) {
	fmt.Printf("下载%s", url)
}
