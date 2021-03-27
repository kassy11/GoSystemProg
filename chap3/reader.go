package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `一行目
二行目
三行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
