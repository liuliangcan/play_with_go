package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, x int
	Fscan(in, &n, &k)
	a := make([]struct{ op, y int }, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].op, &a[i].y)
		if a[i].op == 1 {
			a[i].y -= x // 标记为x的增加量
			x += a[i].y
		} else {
			x += a[i].y
		}
	}
	h := hp{}
	fix1 := 0
	fix2 := 0
	f := 0
	for i := n - 1; i >= 0; i-- {
		if k <= 0 {
			break
		}
		op := a[i].op
		y := a[i].y
		if op == 1 {
			fix1 += y
			k -= 1
		} else {
			if y < 0 {
				fix2 += y
				heap.Push(&h, -y)
			}
		}
		for h.Len() > k {
			fix2 += heap.Pop(&h).(int)
		}
		if fix1+fix2 < f {
			f = fix1 + fix2
		}
	}

	Fprintln(out, x-f)
}

func main() { run(os.Stdin, os.Stdout) }

// 小顶堆
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
