package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var u, v int
	Fscan(in, &u, &v)
	if u > v || (v-u)%2 == 1 {
		Fprintln(out, -1)
		return
	} else if u == v {
		if u == 0 {
			Fprintln(out, "0")
		} else {
			Fprintln(out, "1\n", u)
		}
		return
	}
	x := (v - u) / 2
	if x&u > 0 {
		Fprintln(out, "3\n", u, x, x)
	} else {
		Fprintln(out, "2\n", x, u|x)
	}

}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
