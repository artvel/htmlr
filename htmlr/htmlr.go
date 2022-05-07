package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/artvel/htmlr"
)

func main() {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
		os.Args = remove(os.Args, 1)
	}
	out := flag.String("o", "", "output File")

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage: %s <template.html> -o <output.html>\n", os.Args[0])
		flag.PrintDefaults()
	}

	if input == "" {
		fmt.Println("input template path must be provided")
		flag.Usage()
		return
	}

	output := *out

	if output == "" {
		fmt.Println("output template path must be provided")
		flag.Usage()
		return
	}

	htmlr.Resolve(input, output)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
