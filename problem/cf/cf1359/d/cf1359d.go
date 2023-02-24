package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for mx := 1; mx <= 30; mx++ {
		for i := 0; i < n; {
			for i < n && a[i] > mx {
				i++
			}
			if i == n {
				break
			}
			s := a[i]
			i++
			for i < n && a[i] <= mx {
				s += a[i]
				if s < a[i] {
					s = a[i]
				}
				if s-mx > ans {
					ans = s - mx
				}
				i++
			}
		}
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
