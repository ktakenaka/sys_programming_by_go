package main

import (
	"compress/gzip"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("response.txt.gz")
	gzipwriter := gzip.NewWriter(file)
	defer gzipwriter.Close()
	
	gzipwriter.Header.Name = "request.txt"
	writer := io.MultiWriter(gzipwriter, os.Stdout)

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(writer, conn)
}
