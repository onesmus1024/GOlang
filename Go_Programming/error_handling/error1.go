package main

import (
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}