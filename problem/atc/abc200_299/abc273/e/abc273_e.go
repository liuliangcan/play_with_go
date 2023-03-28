package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var q, x int
	var op string
	Fscan(in, &q)
	type node struct {
		x int
		p *node
	}
	root := &node{x: -1}
	root.p = root
	cur := root
	book := map[int]*node{}
	for i := 0; i < q; i++ {
		Fscan(in, &op)
		if op[0] == 'A' {
			Fscan(in, &x)
			cur = &node{x, cur}
		} else if op[0] == 'D' {
			cur = cur.p
		} else if op[0] == 'S' {
			Fscan(in, &x)
			book[x] = cur
		} else {
			Fscan(in, &x)
			if o, ok := book[x]; ok {
				cur = o
			} else {
				cur = root
			}
		}
		Fprint(out, cur.x, " ")
	}

}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
