package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Example of io.NewReader\n")
	// 読み出し位置と長さを固定したいときに使う
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
