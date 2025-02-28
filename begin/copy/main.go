package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(SrcName, DstName string) {
	src, err := os.Open(SrcName)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(DstName, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer dst.Close()
	_, _ = io.Copy(dst, src)
}
func main() {
	CopyFile("1.txt", "copy.txt")
}
