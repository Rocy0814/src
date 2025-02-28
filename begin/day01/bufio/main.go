package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("read file failed", err)
	}
}
