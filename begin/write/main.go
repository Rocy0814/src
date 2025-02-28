package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("xx.txt", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("open file failed", err)
	}
	defer file.Close()
	file.Write([]byte("Hello Rocy!\n"))
	file.WriteString("你好啊Rocy！")
}
