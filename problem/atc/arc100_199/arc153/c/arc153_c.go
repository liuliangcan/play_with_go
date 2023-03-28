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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	if n == 1 {
		Fprintln(out, "Yes\n0")
		return
	}
	p, s := 0, 0
	x := -int(1e11)
	var b []int
	for i := 0; i < n-1; i++ {
		s += a[i]
		p += a[i] * x
		b = append(b, x)
		x += 1
		if s == a[len(a)-1] && x < 0 {
			x = 0
		}
	}
	if b[len(b)-1] < -p*a[n-1] && -p*a[n-1] <= int(1e12) {
		Fprintln(out, "Yes")
		for _, v := range b {
			Fprint(out, v, " ")
		}
		Fprint(out, -p*a[len(a)-1])
		return
	}
	p, s = 0, 0
	x = int(1e11)
	b = []int{}
	for i := n - 1; i > 0; i-- {
		s += a[i]
		p += a[i] * x
		b = append(b, x)
		x -= 1
		if s == a[0] && x > 0 {
			x = 0
		}
	}
	if -int(1e12) < -p*a[0] && -p*a[0] < b[len(b)-1] {
		Fprint(out, "Yes\n", -p*a[0], " ")
		for i := len(b) - 1; i >= 0; i-- {
			Fprint(out, b[i], " ")
		}
		return
	}

	Fprintln(out, "No")
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
