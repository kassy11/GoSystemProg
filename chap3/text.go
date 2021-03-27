package main

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	file, err := os.Open("lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks := readChunks(file)
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	crc.Write(byteData)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}
