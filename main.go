package main

import (
	"flag"
	"fmt"

	"./randpack"
	"./text"
)

func main() {
	flag.Parse()
	fp := flag.Args()
	texts := text.ReadText(fp[0])

	r := randpack.Random{Param: len(texts)}
	fmt.Println(texts[r.Roulette()])
	//fmt.Println(r.Roulette())
}
