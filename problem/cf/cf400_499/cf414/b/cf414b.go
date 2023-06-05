package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const MOD = int(1e9 + 7)

func solve(in io.Reader, out io.Writer) {
	var n, u, ans int
	Fscan(in, &u, &n)
	f := make([]int, u+1)
	for i := 1; i <= u; i++ {
		f[i] = 1
	}
	for i := 1; i < n; i++ {
		g := make([]int, u+1)
		for j := 1; j <= u; j++ {
			for k := j; k <= u; k += j {
				g[k] = (g[k] + f[j]) % MOD
			}
		}
		f = g
	}
	for _, v := range f {
		ans = (ans + v) % MOD
	}

	Fprintln(out, ans)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
