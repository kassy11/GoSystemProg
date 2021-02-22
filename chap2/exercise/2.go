package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	file, err := os.Create("2.csv")
	if err != nil {
		panic(err)
	}
	source := [][]string{
		{"hello", "world"},
		{"kotaro", "kashihara"},
		{"asuka", "saito"},
	}
	WriteCsv(file, source)
}

func WriteCsv(output io.Writer, source [][]string) {
	w := csv.NewWriter(output)
	for _, line := range source {
		if err := w.Write(line); err != nil {
			panic(err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		panic(err)
	}
}
