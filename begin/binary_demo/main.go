package main

import "fmt"

const (
	eat   int = 4 //100
	sleep int = 2 //010
	da    int = 1 //001
)

func f(arg int) {
	fmt.Printf("%b\n", arg)
}
func main() {
	f(eat | sleep)
	f(eat | sleep | da)
}
