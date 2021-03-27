package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("filename needed")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(newfile, file)
}
