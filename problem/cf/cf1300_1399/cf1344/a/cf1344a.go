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
	a := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a[((v+i)%n+n)%n] = 1
	}
	//print(len(a))
	if len(a) == n {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}

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
