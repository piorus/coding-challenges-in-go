package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(f)

	i := 0

	for scanner.Scan() {
		i = i + 1
	}

	fmt.Println(i, os.Args[1])
}
