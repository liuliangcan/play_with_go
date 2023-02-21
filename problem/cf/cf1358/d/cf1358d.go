package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, x, s, p, ans int
	Fscan(in, &n, &x)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &d[i])
	}
	d = append(d, d...) // 首尾相接环形数组变成普通数组滑窗
	type pair struct{ l, r int }
	var q []pair // 用队列模拟滑窗
	for _, v := range d {
		q = append(q, pair{1, v}) // 由于最优方案一定使滑窗末尾在某月末尾，因此直接滑窗纳入本月
		p += v
		s += (1 + v) * v / 2
		for p > x { // 滑窗里元素超了则从左边尝试减去对应天数
			l, r := q[0].l, q[0].r
			if r-l+1 <= p-x { // 减去窗左边整个月都不够
				s -= (l + r) * (r - l + 1) / 2
				p -= r - l + 1
				q = q[1:]
			} else {
				c := p - x
				s -= (l + l + c - 1) * c / 2
				p -= c
				q[0].l = l + c
			}
		}
		if ans < s {
			ans = s
		}
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
