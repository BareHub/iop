package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	s := strings.NewReader("1234567890123456789012345678901234567890123456789012345678901234567890")
	r := bufio.NewReaderSize(s, 16)

	i := 0
	for {
		i++
		buf := make([]byte, 6)
		var (
			size int
			err  error
		)
		if size, err = io.ReadFull(r, buf); err != nil {
			log.Fatalf("[%3d|%d] %q --- error: %v\n", i, size, buf, err)
		}
		fmt.Printf("[%3d|%d] %q\n", i, size, buf)
	}
}
