package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("strings builder example\n"))
	// 内部では unsafe を使っているので、allocationせずにbyteの配列をstringに変換している
	fmt.Println(builder.String())
}
