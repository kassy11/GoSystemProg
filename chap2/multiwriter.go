package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	// 複数のio.Writerを受け取り、それら全てに対して同時に書き込む
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
