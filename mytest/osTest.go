package main

import (
	"os"
	"fmt"
)

func main() {
	num1 := os.Getpagesize()
	fmt.Println(num1)
}
