package main

import (
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
	

	fmt.Println(os.Args)
}
