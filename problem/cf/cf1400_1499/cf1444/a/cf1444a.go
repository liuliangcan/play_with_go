package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var p, q int
	Fscan(in, &p, &q)
	if p%q != 0 {
		Fprintln(out, p)
		return
	}
	ans := 1
	for k, v := range getPrimeReasons(q) {
		cnt := v
		x := p / powint(k, v)
		for x%k == 0 {
			cnt += 1
			x /= k
		}
		ans = max(ans, x*powint(k, v-1))
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

// / getPrimeReasons
// / 分解质因数模板
func getPrimeReasons(x int) map[int]int {
	ans := map[int]int{}
	if x == 1 {
		return ans
	}
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			ans[i]++
			x /= i
		}
	}
	if x > 1 {
		ans[x]++
	}
	return ans
}
func powint(a, e int) int {
	res, m := 1, a
	for e > 0 {
		if e&1 != 0 {
			res = res * m
		}
		m = m * m
		e >>= 1
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
