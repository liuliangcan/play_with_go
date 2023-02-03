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
	var n, ans, s, v int
	cnt := map[int]int{0: 1}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		s += v * (n%2*2 - 1)
		ans += cnt[s]
		cnt[s] += 1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
