package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "%d歳の%sさんが%f秒で到着した\n", 20, "和田", 3.14)
}
