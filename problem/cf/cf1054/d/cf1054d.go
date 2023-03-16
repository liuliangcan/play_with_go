package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, s, v, k int
	Fscan(in, &n, &k)
	mask := (1 << k) - 1
	cnt := map[int]int{0: 1}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s ^= v
		cnt[min(s, s^mask)]++
	}
	ans := c2(n + 1)
	for _, v := range cnt {
		ans -= c2(v/2) + c2(v-v/2)
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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func c2(x int) int64 {
	return int64(x-1) * int64(x) / 2
}
