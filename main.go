package main

import (
	"io"
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

func kclipCopy(dst io.Writer, src io.Reader) (int64, error) {
	pads := padList.Current()
	head := pads.Head
	tail := pads.Tail
	return copyPad(dst, src, head, tail)
}

func copyPad(dst io.Writer, src io.Reader, head, tail string) (written int64, err error) {
	nw, err := io.WriteString(dst, head)
	if err != nil {
		return
	}
	written += int64(nw)
	wtmp, err := io.Copy(dst, src)
	written += wtmp
	if err != nil {
		return
	}
	nw, err = io.WriteString(dst, tail)
	written += int64(nw)
	return
}
