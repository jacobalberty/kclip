package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	clipboard := cList.Current(os.Stdout)
	if len(os.Args) == 1 {
		_, err := io.Copy(clipboard, os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		for _, fname := range os.Args[1:] {
			fh, err := os.Open(fname)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(clipboard, fh)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("Sending to clipboard via %s control codes\n", strings.Join(clipboard.SList, ","))
}
