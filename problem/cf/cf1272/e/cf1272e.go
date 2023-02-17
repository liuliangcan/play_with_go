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
	var n, v int
	Fscan(in, &n)
	a := make([]int, n+1)
	g := make([][]int, n+1)
	var q0, q1 []int
	d0, d1 := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		for _, v = range []int{i - a[i], i + a[i]} {
			if 1 <= v && v <= n {
				g[v] = append(g[v], i)
			}
		}
		if a[i]%2 == 0 {
			q0 = append(q0, i)
			d0[i], d1[i] = 0, -1
		} else {
			q1 = append(q1, i)
			d1[i], d0[i] = 0, -1
		}
	}
	bfs := func(q []int, dist *[]int) {
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v = range g[u] {
				if (*dist)[v] == -1 {
					(*dist)[v] = (*dist)[u] + 1
					q = append(q, v)
				}
			}
		}
	}
	bfs(q0, &d0)
	bfs(q1, &d1)
	//bfs := func(q []int, dist []int) {
	//	for len(q) > 0 {
	//		u := q[0]
	//		q = q[1:]
	//		for _, v = range g[u] {
	//			if dist[v] == -1 {
	//				dist[v] = dist[u] + 1
	//				q = append(q, v)
	//			}
	//		}
	//	}
	//}
	//bfs(q0, d0)
	//bfs(q1, d1)
	for i := 1; i <= n; i++ {
		if a[i]%2 == 0 {
			Fprint(out, d1[i], " ")
		} else {
			Fprint(out, d0[i], " ")
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
