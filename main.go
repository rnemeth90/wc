package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	lines bool
	bites bool // clever
	files string
)

func init() {
	flag.BoolVar(&lines, "l", false, "count lines")
	flag.BoolVar(&bites, "b", false, "count bytes")
	flag.StringVar(&files, "f", "", "a space delimited list of files")
}

func main() {
	flag.Parse()

	if files != "" {
		count, err := readFile(files, lines, bites)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(count)
		os.Exit(0)
	}

	count, err := count(os.Stdin, lines, bites)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(count)
}

func readFile(files string, countLines bool, countBytes bool) (int, error) {
	fs := strings.Split(files, " ")
	var c int

	for _, f := range fs {
		file, err := os.Open(f)
		if err != nil {
			return 0, err
		}
		defer file.Close()
		c, err = count(file, countLines, countBytes)
	}
	return c, nil
}

func count(r io.Reader, countLines bool, countBytes bool) (int, error) {
	scanner := bufio.NewScanner(r)

	if countBytes && countLines {
		return 0, errors.New("You cannot count lines and bytes")
	} else if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc, nil
}
