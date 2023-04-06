package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	var str string
	Fscan(in, &n)
	type xs struct{ x, s int }
	a := make([]xs, n)
	for i := 0; i < n; i++ {
		Fscan(in, &str)
		for _, c := range str {
			if c == 'X' {
				a[i].x++
			} else {
				a[i].s += int(c - '0')
				ans += a[i].x * int(c-'0')
			}
		}
	}
	//println(ans)
	sort.Slice(a, func(i, j int) bool { return a[i].x*a[j].s < a[i].s*a[j].x })
	p := 0
	for _, z := range a {
		//println(z.x, z.s)
		ans += z.x * p
		p += z.s
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
