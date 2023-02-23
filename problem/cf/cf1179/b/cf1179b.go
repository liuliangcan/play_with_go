package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	l, r := 0, n*m-1
	for l <= r {
		Fprintln(out, l/m+1, l%m+1)
		if l < r {
			Fprintln(out, r/m+1, r%m+1)
		} else if l == r {
			break
		}
		l++
		r--
	}
}

func main() { run(os.Stdin, os.Stdout) }
