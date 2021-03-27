package main

import (
	"fmt"
	"strings"
)

var num_source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(num_source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v g=%#v s=%#v\n", i, f, g, s)
}
