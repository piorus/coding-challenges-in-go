package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	linesPtr := flag.Bool("l", false, "count lines")
	bytesPtr := flag.Bool("c", false, "count bytes")
	flag.Parse()

	f, err := os.Open(flag.Args()[0])
	check(err)
	defer f.Close()

	if *linesPtr {
		scanner := bufio.NewScanner(f)
		lines := 0

		for scanner.Scan() {
			lines = lines + 1
		}

		fmt.Println("lines", lines)
	}

	if *bytesPtr {
		fs, err := f.Stat()
		check(err)
		fmt.Println("bytes", fs.Size())
	}
}
