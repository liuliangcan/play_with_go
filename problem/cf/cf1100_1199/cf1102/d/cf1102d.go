package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	var s []byte
	Fscan(in, &n, &s)
	cnt := make([]int, 3)
	for _, c := range s {
		cnt[c-'0']++
	}
	k := n / 3
	for i, c := range s {
		c -= '0'
		if cnt[c] > k {
			for j := byte(0); j < c; j++ {
				if cnt[j] < k {
					cnt[j]++
					cnt[c]--
					s[i] = j + '0'
					break
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		c := s[i] - '0'
		if cnt[c] > k {
			for j := byte(2); j > c; j-- {
				if cnt[j] < k {
					cnt[j]++
					cnt[c]--
					s[i] = j + '0'
					break
				}
			}
		}
	}

	Fprintf(out, "%s", s)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
