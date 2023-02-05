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

	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := n - 2; i >= 0; i-- {
		t1, t2 := a[i], a[i+1]
		if t1 <= t2 {
			continue
		}
		l := len(t1)
		if len(t2) < l {
			l = len(t2)
		}
		if t1[:l] == t2[:l] {
			a[i] = t1[:l]
			continue
		}

		for j := 0; j < l; j++ {
			if t1[j] > t2[j] {
				a[i] = t1[:j]
				break
			}
		}
	}
	for _, v := range a {
		Fprintln(out, v)
	}

}

func main() { run(os.Stdin, os.Stdout) }
