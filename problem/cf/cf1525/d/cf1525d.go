package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	var chairs, persons []int
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v == 1 {
			persons = append(persons, i)
		} else {
			chairs = append(chairs, i)
		}
	}
	n = len(persons)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = 1e9
	}
	for i, v := range chairs {
		for j := min(i, n-1); j >= 0; j-- {
			f[j+1] = min(f[j+1], f[j]+abs(v-persons[j]))
		}
	}

	Fprintln(out, f[n])
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
