package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("filename is missing")
	}

	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	linesPtr := flag.Bool("l", false, "count lines")
	bytesPtr := flag.Bool("c", false, "count bytes")

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
