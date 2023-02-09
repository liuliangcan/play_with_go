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
	Fscan(in, &n)
	a := make([]int, n+1)
	type pair struct{ l, r int }
	d := map[int][]pair{}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		for j, s := i, 0; j > 0; j-- {
			s += a[j]
			d[s] = append(d[s], pair{j, i})
		}
	}
	var ans []pair
	for _, p := range d {
		if len(p) <= len(ans) { // 总长度都达不到当前ans，一定没用
			continue
		}
		var t []pair
		r0 := 0
		for _, v := range p {
			if v.l > r0 {
				t = append(t, pair{v.l, v.r})
				r0 = v.r
			}
		}
		if len(t) > len(ans) {
			ans = t
		}
	}

	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p.l, p.r)
	}

}

func main() { run(os.Stdin, os.Stdout) }
