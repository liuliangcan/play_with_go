package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	. "strings"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, k, x int
	Fscan(in, &k, &x, &n, &m)
	for c1 := 0; c1 < 2; c1++ {
		for a1 := 0; a1 < 2; a1++ {
			if a1+c1 > n {
				continue
			}
			for ac1 := 0; ac1 <= (n-a1-c1)/2; ac1++ {
				for c2 := 0; c2 < 2; c2++ {
					for a2 := 0; a2 < 2; a2++ {
						if a2+c2 > m {
							continue
						}
						for ac2 := 0; ac2 <= (m-a2-c2)/2; ac2++ {
							C1, AC1, A1, C2, AC2, A2 := c1, ac1, a1, c2, ac2, a2
							for i := 2; i < k && AC2 <= x; i++ {
								C1, AC1, A1, C2, AC2, A2 = C2, AC2, A2, C1, AC1+AC2+(A1&C2), A2
							}
							if AC2 == x {
								Fprintln(out, Repeat("C", c1)+Repeat("AC", ac1)+Repeat("Z", n-c1-a1-ac1*2)+Repeat("A", a1))
								Fprintln(out, Repeat("C", c2)+Repeat("AC", ac2)+Repeat("Z", m-c2-a2-ac2*2)+Repeat("A", a2))
								//Fprintln(out, []string{"", "C"}[c1]+Repeat("AC", ac1)+Repeat("Z", n-c1-a1-ac1*2)+[]string{"", "A"}[a1])
								//Fprintln(out, []string{"", "C"}[c2]+Repeat("AC", ac2)+Repeat("Z", m-c2-a2-ac2*2)+[]string{"", "A"}[A1])
								return
							}
						}
					}
				}
			}
		}
	}

	Fprintln(out, "Happy new year!")
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
