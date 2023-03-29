package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	if n == 1 {
		Fprintln(out, -1)
		return
	}
	sort.Ints(a)
	if a[0] == a[len(a)-1] {
		Fprintln(out, "1\n", a[0])
		return
	}
	var ans []int
	if n == 2 {
		d := a[1] - a[0]
		ans = append(ans, a[0]-d)
		if d&1 == 0 {
			ans = append(ans, a[0]+d/2)
		}
		ans = append(ans, a[1]+d)
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		return
	}
	cnt := map[int]int{}
	for i := 0; i < n-1; i++ {
		cnt[a[i+1]-a[i]]++
		if len(cnt) >= 3 {
			Fprintln(out, 0)
			return
		}
	}
	if len(cnt) == 1 {
		for d := range cnt {
			Fprintln(out, "2\n", a[0]-d, a[len(a)-1]+d)
			return
		}
	}
	if len(cnt) == 2 {
		for d := range cnt {
			ans = append(ans, d)
		}
		sort.Ints(ans)
		if cnt[ans[1]] != 1 || ans[1] != ans[0]*2 {
			Fprintln(out, 0)
			return
		}
		Fprintln(out, 1)
		for i := 0; i < n-1; i++ {
			if a[i+1]-a[i] == ans[1] {
				Fprintln(out, a[i]+ans[0])
				return
			}
		}
	}

	Fprintln(out, 0)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
