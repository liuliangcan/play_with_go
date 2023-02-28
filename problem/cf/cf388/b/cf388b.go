package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var k, s int
	Fscan(in, &k)
	var a []int
	for k > 0 {
		if k < 32 {
			break
		}
		x := int(math.Pow(float64(k), 0.2))
		k -= x * x * x * x * x
		a = append(a, x)
		s += x
	}
	n := 2 + s*4 + 4 + k
	start := 2
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, n)
		for j := range ans[i] {
			ans[i][j] = 'N'
		}
	}
	var add = func(x, y int) {
		ans[x][y], ans[y][x] = 'Y', 'Y'
	}
	for _, x := range a {
		for i := 0; i < 4; i++ {
			for y := 0; y < x; y++ {
				l := start + i*x + y
				for z := 0; z < x; z++ {
					r := start + (i+1)*x + z
					add(l, r)
				}
			}
		}
		for y := 0; y < x; y++ {
			l := start + y
			r := start + 4*x + y
			add(0, l)
			add(r, 0)
		}
		start += 5 * x
	}
	for i := start; i < start+k; i++ {
		add(0, i)
		add(i, start+k)
	}
	start += k
	for i := 0; i < 3; i++ {
		add(start, start+1)
		start++
	}
	add(start, 1)

	Fprintln(out, n)
	for i := 0; i < n; i++ {
		Fprintln(out, string(ans[i]))
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
