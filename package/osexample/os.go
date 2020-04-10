package osexample

import (
	"fmt"
	"log"
	"os"
)

// OpenFile example
func OpenFile() {

	f, err := os.OpenFile("hello.txt", os.O_CREATE, os.ModeType)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	n, err := f.WriteString("hello world1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
