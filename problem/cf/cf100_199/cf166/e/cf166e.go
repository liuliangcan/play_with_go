package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	abc, d := 0, 1
	const MOD = int(1e9) + 7
	for i := 0; i < n; i++ {
		abc, d = (d+abc*2)%MOD, abc*3%MOD
	}

	Fprintln(out, d)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
