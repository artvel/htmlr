package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/artvel/htmlr"
)

func main() {
	// verbose := flag.Bool("v", false, "verbose")
	output := flag.String("o", "", "output File")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("Usage: %s <template.html> -o <output.html>\n", os.Args[0])
		flag.PrintDefaults()
	}
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	if input == "" {
		flag.Usage()
		return
	}

	htmlr.Resolve(input, *output)
}
