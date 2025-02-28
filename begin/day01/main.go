package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Println("open file failed!", err)
	// 	return
	// }
	// defer file.Close()
	// var tmp = make([]byte, 128)
	// n, err := file.Read(tmp)
	// if err == io.EOF {
	// 	fmt.Println("文件读完了")
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("read file errored", err)
	// }
	// fmt.Printf("读取了%d字节的数据\n", n)
	// fmt.Println(string(tmp[:n]))
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed", err)
	}
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
