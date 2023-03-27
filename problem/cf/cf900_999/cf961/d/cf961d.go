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
	type pt struct{ x, y int }
	ps := make([]pt, n)
	for i := 0; i < n; i++ {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	if n <= 4 {
		Fprintln(out, "YES")
		return
	}
	is_line := func(a, b, c pt) bool { // 判断三点是否共线
		return (a.y-b.y)*(a.x-c.x) == (a.y-c.y)*(a.x-b.x)
	}
	f := func(a, b pt) bool { // 判断第一条线画在ab是否可行
		var other []pt
		for _, c := range ps {
			if !is_line(a, b, c) {
				if len(other) < 2 {
					other = append(other, c)
				} else if !is_line(other[0], other[1], c) {
					return false
				}
			}
		}
		return true
	}
	if f(ps[0], ps[1]) || f(ps[0], ps[2]) || f(ps[2], ps[1]) {
		Fprintln(out, "YES")
		return
	}
	Fprintln(out, "NO")

}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
