package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "My name is Debajit!"
	r := strings.NewReader(str)

	// Creating abyte slice of 8 byte which will be filled up while reading the string using Reader
	bs := make([]byte, 8)
	for {
		n, err := r.Read(bs)
		fmt.Printf("n = %v, err = %v, bs = %v\n", n, err, bs)
		fmt.Printf("Content: %q\n", bs[:n])

		if err!= nil && err == io.EOF {
			break
		}
	}
}