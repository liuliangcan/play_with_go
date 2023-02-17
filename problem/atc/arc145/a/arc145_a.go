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
	var n int
	var s string
	Fscan(in, &n, &s)
	if s == "BA" || s[0] == 'A' && s[n-1] == 'B' {
		Fprintf(out, "No")
	} else {
		Fprintf(out, "Yes")
	}
}

func main() { run(os.Stdin, os.Stdout) }
