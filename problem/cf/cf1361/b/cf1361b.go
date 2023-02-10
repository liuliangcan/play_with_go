package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

const MOD int64 = 1e9 + 7

func modPow(a, b, mod int64) int64 {
	var ans int64 = 1
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			ans = ans * a % mod
		}
		a = a * a % mod
	}
	return ans
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, t, p int

	type pair struct{ k, c int }
o:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &p)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if p == 1 {
			Fprintln(out, n%2)
			continue o
		}
		sort.Ints(a)
		targetK := a[n-1]
		var st []pair
		for i := n - 2; i >= 0; i-- {
			k := a[i]
			for len(st) > 0 && st[len(st)-1] == (pair{k, p - 1}) {
				st = st[:len(st)-1]
				k++
			}
			if k == targetK {
				if i == 0 {
					Fprintln(out, 0)
					continue o
				}
				i--
				targetK = a[i]
			} else if len(st) > 0 && st[len(st)-1].k == k {
				st[len(st)-1].c++
			} else {
				st = append(st, pair{k, 1})
			}
		}
		ans := modPow(int64(p), int64(targetK), MOD)
		for _, v := range st {
			ans = (ans - modPow(int64(p), int64(v.k), MOD)*int64(v.c)) % MOD
		}

		Fprintln(out, (ans+MOD)%MOD)
	}

}

func main() { run(os.Stdin, os.Stdout) }
