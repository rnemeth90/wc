package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	lineCount bool
	byteCount bool
	file      string
)

func init() {
	flag.BoolVar(&lineCount, "l", false, "Count lines")
	flag.BoolVar(&byteCount, "b", false, "Count Bytes")
	flag.StringVar(&file, "f", "", "[optional] file")
	flag.Usage = usage
}

func usage() {
	fmt.Println(os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	standardArgs := flag.Args()

	// reading from stdin if -f is not passed and nothing on os.args
	if file == "" && len(standardArgs) == 0 {
		fmt.Println(count(os.Stdin, byteCount, lineCount))
		os.Exit(0)
		// get the file from the -f flag
	} else if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(count(f, byteCount, lineCount))
		os.Exit(0)
		// get multiple files from os.Args
	} else if len(standardArgs) > 0 {
		for _, file := range standardArgs {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
			count := count(f, byteCount, lineCount)
			fmt.Println(count)
		}
		os.Exit(0)
	}
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
