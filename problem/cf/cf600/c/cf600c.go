package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s string
	Fscan(in, &s)
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}
	l, r := 0, 25
	for l < r {
		for l < r && cnt[l]%2 == 0 {
			l++
		}
		for l < r && cnt[r]%2 == 0 {
			r--
		}
		if l < r {
			cnt[l]++
			cnt[r]--
			l++
			r--
		}
	}
	mid := ""
	for c, v := range cnt {
		Fprint(out, strings.Repeat(string(byte('a'+c)), v/2))
		if v%2 == 1 {
			mid = string(byte('a' + c))
		}
	}
	Fprint(out, mid)
	for i := 25; i >= 0; i-- {
		Fprint(out, strings.Repeat(string(byte('a'+i)), cnt[i]/2))
	}
}

func main() { run(os.Stdin, os.Stdout) }
