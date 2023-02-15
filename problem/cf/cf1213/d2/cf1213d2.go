package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// 233 ms
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 2e5
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	p := make([][]int, mx+1)
	for _, v := range a {
		i := 0
		for v > 0 {
			p[v] = append(p[v], i)
			i++
			v >>= 1
		}
		p[v] = append(p[v], i)
	}
	ans := int(1e9)
	for _, b := range p {
		if len(b) >= k {
			t := 0
			for i := 0; i < k; i++ {
				t += b[i]
			}
			ans = min(ans, t)
		}
	}

	Fprintln(out, ans)
}

// 139 ms
func run1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 2e5
	p := make([][]int, mx+1)
	var n, k, v int
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &v)
		if p[v] == nil {
			p[v] = []int{1}
		} else {
			p[v][0]++
		}
	}
	ans := int(1e9)
	for i := mx; i > 0; i-- {
		b := p[i]
		if b == nil {
			continue
		}
		s, left := 0, k
		for j, c := range b {
			if left <= c {
				ans = min(ans, s+left*j)
				break
			}
			s += c * j
			left -= c
		}
		i2 := i >> 1
		if p[i2] == nil {
			p[i2] = append([]int{0}, b...)
		} else {
			for j, c := range b {
				if j+1 == len(p[i2]) {
					p[i2] = append(p[i2], b[j:]...)
					break
				}
				p[i2][j+1] += c
			}
		}
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
