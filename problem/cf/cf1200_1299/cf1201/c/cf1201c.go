package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	a = a[n/2:]
	i := 1
	n = len(a)
	for ; i < n; i++ {
		if (a[i]-a[i-1])*i > k {
			break
		}
		k -= (a[i] - a[i-1]) * i
	}

	Fprintln(out, a[i-1]+k/i)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
