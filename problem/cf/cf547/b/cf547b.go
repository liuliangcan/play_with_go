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
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	l := make([]int, n)
	r := make([]int, n)
	var st1, st2 []int
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		l[i] = -1
		r[i] = n
		for len(st1) > 0 && a[st1[len(st1)-1]] >= a[i] {
			st1 = st1[:len(st1)-1]
		}
		if len(st1) > 0 {
			l[i] = st1[len(st1)-1]
		}
		st1 = append(st1, i)
		for len(st2) > 0 && a[st2[len(st2)-1]] > a[i] {
			r[st2[len(st2)-1]] = i
			st2 = st2[:len(st2)-1]
		}
		st2 = append(st2, i)
	}
	ans := make([]int, n)
	for i, v := range a {
		x := r[i] - l[i] - 2
		if v > ans[x] {
			ans[x] = v
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ans[i] < ans[i+1] {
			ans[i] = ans[i+1]
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
