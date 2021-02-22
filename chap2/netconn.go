package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	// net.Connインタフェースが返ってくる
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	// net.Connはio.Writerインタフェースを実装している
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// io.Readerインタフェースも実装している
	io.Copy(os.Stdout, conn)

	fmt.Println("-------------------------------------")

	// HTTPリクエストの作成に使う構造体を利用して、リクエストを送る
	req, err := http.NewRequest("GET", "http://ascii.jp:80", nil)
	req.Write(conn)
	req.Write(os.Stdout)
}
