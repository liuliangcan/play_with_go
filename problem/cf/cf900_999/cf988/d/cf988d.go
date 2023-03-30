package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	var ans []int
	mx := int(-1e10)
	a := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a[v] = 1
		if v > mx {
			mx = v
		}
	}
	ans = []int{mx}
	for x := range a {
		for k := 0; k <= 32; k++ {
			t := x + (1 << k)
			if t > mx {
				break
			}
			if a[t] == 1 {
				t2 := x + (1 << (k + 1))
				if a[t2] == 1 {
					Fprintln(out, "3\n", x, t, t2)
					return
				}
				ans = []int{x, t}
			}
		}
	}

	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
