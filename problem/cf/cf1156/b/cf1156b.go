package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	g := make([][]byte, 2)
	for _, c := range s {
		g[c&1] = append(g[c&1], c)
	}

	x, y := g[0], g[1]
	if len(x) == 0 {
		Fprintf(out, "%s\n", y)
	} else if len(y) == 0 {
		Fprintf(out, "%s\n", x)
	} else if 1 != abs(int(x[len(x)-1])-int(y[0])) {
		Fprintf(out, "%s%s\n", x, y)
	} else if 1 != abs(int(y[len(y)-1])-int(x[0])) {
		Fprintf(out, "%s%s\n", y, x)
	} else {
		Fprintln(out, "No answer")
	}

}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t int
	for Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
