package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	// io.WriterインタフェースのWriteのポインタ
	writer := gzip.NewWriter(file)
	writer.Header.Name = "text.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}
