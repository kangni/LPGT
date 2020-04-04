package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "hello", ""
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i)
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
