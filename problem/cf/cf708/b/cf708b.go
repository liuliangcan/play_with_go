package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"strings"
)

func solve(in io.Reader, out io.Writer) {
	var a, b, c, d int
	Fscan(in, &a, &b, &c, &d)
	if a+b+c+d == 0 {
		Fprintf(out, "0")
		return
	}
	zeros := []int{int((1 + math.Sqrt(float64(1+8*a))) / 2), int((1 - math.Sqrt(float64(1+8*a))) / 2)}
	ones := []int{int((1 + math.Sqrt(float64(1+8*d))) / 2), int((1 - math.Sqrt(float64(1+8*d))) / 2)}
	for _, zero := range zeros {
		for _, one := range ones {
			if zero < 0 || one < 0 {
				continue
			}
			if zero*(zero-1)/2 != a || one*(one-1)/2 != d || (zero+one)*(zero+one-1)/2 != a+b+c+d {
				continue
			}
			if zero == 0 {
				Fprintln(out, strings.Repeat("1", one))
				return
			}
			if one == 0 {
				Fprintln(out, strings.Repeat("0", zero))
				return
			}
			back1 := b / zero
			mod := b % zero
			front1 := one - back1
			var front []byte
			var zs []byte
			for i := 0; i < zero; i++ {
				zs = append(zs, '0')
			}
			for i := 0; i < front1; i++ {
				front = append(front, '1')
			}
			if mod > 0 {
				front[len(front)-1] = '0'
				zs[mod-1] = '1'
			}
			Fprint(out, string(front))
			Fprint(out, string(zs))
			Fprintln(out, strings.Repeat("1", back1))
			return
		}
	}
	Fprintln(out, "Impossible")
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
