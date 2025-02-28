package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.OpenFile("./XX.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
}
