package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	file.Write([]byte("os.Create example\n"))
	file.Close()
}
