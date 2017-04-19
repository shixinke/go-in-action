package main

import (
	"fmt"
	"spider"
)

func init() {
	fmt.Println("caller init")
}

func main() {
	spider.Run()
}
