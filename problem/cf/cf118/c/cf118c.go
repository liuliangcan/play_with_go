package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, k int
	var s, ans string
	Fscan(in, &n, &k, &s)
	poses := ['9' + 1][]int{}
	for i, c := range s {
		poses[c] = append(poses[c], i)
	}

	mn := int(1e9)
	for i := byte('0'); i <= byte('9'); i++ {
		cost := 0
		t := []byte(s)
		remain := k - len(poses[i])
		for dis := byte(1); dis <= byte(9); dis++ {
			j := i + dis
			if j <= '9' {
				ps := poses[j]
				for c := 0; c < len(ps) && remain > 0; c++ {
					t[ps[c]] = i
					remain--
					cost += int(dis)
				}
			}
			j = i - dis
			if j >= '0' {
				ps := poses[j]
				for c := len(ps) - 1; c >= 0 && remain > 0; c-- {
					t[ps[c]] = i
					remain--
					cost += int(dis)
				}
			}
			if remain <= 0 || cost > mn {
				break
			}
		}
		if remain <= 0 {
			res := string(t)
			if cost < mn {
				mn = cost
				ans = res
			} else if cost == mn && res < ans {
				ans = res
			}
		}
	}

	Fprintln(out, mn)
	Fprintln(out, ans)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
