package web

import "fmt"

func init() {
	fmt.Println("web init")
}

func GetUrl() string {
	return "https://golang.org"
}
