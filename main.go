package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 {
		_, err := kclipCopy(os.Stdout, os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		for _, fname := range os.Args[1:] {
			fh, err := os.Open(fname)
			if err != nil {
				panic(err)
			}

			_, err = kclipCopy(os.Stdout, fh)
			if err != nil {
				panic(err)
			}
		}
	}
}
