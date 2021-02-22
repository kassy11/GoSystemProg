package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with OS.stdout at %v", time.Now())
}
