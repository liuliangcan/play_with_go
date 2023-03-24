package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, s int
	Fscan(in, &n)
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		cnt[s]++
	}
	s = 0
	mx := map[int]int{} // 每个组里只要最大那个
	for _, v := range cnt {
		s += v
		l := bits.Len(uint(s))
		if s&(s-1) == 0 {
			mx[l-1] = s
		} else if mx[l] == 0 || mx[l] < s {
			mx[l] = s
		}
	}
	var vs []int
	for _, v := range mx {
		vs = append(vs, v)
	}
	sort.Ints(vs)
	f := func(x int) int {
		if x == 0 {
			return 1
		}
		if x&(x-1) == 0 {
			return 0
		}
		t := 1 << bits.Len(uint(x))
		return t - x
	}
	ans := f(0)*2 + f(s) // 全分1组
	suf := 0
	for y := n; y >= 0; y-- {
		suf += cnt[y]
		if cnt[y] > 0 {
			ans = min(ans, f(suf)+f(s-suf)+f(0)) // 分2组
			for _, x := range vs {
				if x+suf >= s { // 分不了3组就跳出
					break
				}
				ans = min(ans, f(x)+f(suf)+f(s-x-suf)) // 分3组
			}
		}
	}
	Fprintln(out, ans)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t int
	for Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
