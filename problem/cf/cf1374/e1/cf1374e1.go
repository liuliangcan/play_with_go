package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, x, y, t, ans int
	Fscan(in, &n, &k)
	var a []int
	var b []int
	var c []int
	for i := 0; i < n; i++ {
		Fscan(in, &t, &x, &y)
		if x == 1 && y == 1 {
			c = append(c, t)
		} else if x == 1 {
			a = append(a, t)
		} else if y == 1 {
			b = append(b, t)
		}
	}

	if len(b) < len(a) {
		a, b = b, a
	}
	if len(c)+len(a) < k {
		Fprintln(out, -1)
		return
	}
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	for i := range a {
		a[i] += b[i] // 合并ab
	}
	both := len(c)
	if both > k {
		both = k
	}
	for i := 0; i < both; i++ {
		ans += c[i] // 先从c中取k个
	}
	j := k - both // a的遍历起点
	for i := 0; i < j; i++ {
		ans += a[i] // 如果c不够，从a中取
	}

	for i := both - 1; i >= 0 && j < len(a) && a[j] < c[i]; j++ {
		ans += a[j] - c[i]
		i--
	}

	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
