package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)

	fi, err := os.Lstat(path)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name %s\n", fi.Name())
	fmt.Printf("Mode %s\n", fi.Mode())
	fmt.Printf("Size %d\n", fi.Size())
	fmt.Printf("IsDir %t\n", fi.IsDir())
	fmt.Printf("ModTime %s\n", fi.ModTime())
}
