package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("I/O test")
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Println(in.Text())
		if in.Text() == "exit" {
			break
		}
	}
}
