package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count Bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *bytes, *lines))
}

func count(r io.Reader, bytes, lines bool) int {
	scanner := bufio.NewScanner(r)
	if !lines && !bytes {
		scanner.Split(bufio.ScanWords)
	} else if bytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanLines)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
