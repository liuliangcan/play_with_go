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
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)
	cnt := 0 // 能完整吃cnt圈
	for _, v := range b {
		d := v - cnt // 当前堆超过cnt个苹果
		if d > 0 {
			if k > d*n { // 剩余n堆苹果多吃这么多,且还能继续
				k -= d * n
				cnt = v
			} else {
				cnt += k / n // 剩余n堆苹果不够吃d圈，只能吃k//n圈
				k %= n
				break
			}
		}
		n--
	}
	for _, v := range a {
		v -= cnt
		if v < 0 {
			v = 0
		}
		if v > 0 && k > 0 {
			k--
			v--
		}
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
