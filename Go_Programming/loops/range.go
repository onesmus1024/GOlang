package main

import (
	"fmt"
	"os"
)


func main() {
// range loop	
	for i := range os.Args {
		fmt.Println(os.Args[i])
	}
}