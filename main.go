package main

import (
	"flag"
	"fmt"

	"./text"
)

func main() {
	flag.Parse()
	fp := flag.Args()
	fmt.Println(fp[0])
	texts := text.ReadText(fp[0])
	fmt.Println(texts)
}
