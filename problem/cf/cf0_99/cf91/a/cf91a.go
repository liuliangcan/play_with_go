package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	var s, t string
	Fscan(in, &s, &t)
	n = len(s)
	dp := make([][26]int, n+1)
	for i := 0; i < 26; i++ {
		dp[n][i] = n
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			dp[i][j] = dp[i+1][j]
		}
		dp[i][s[i]-'a'] = i
	}
	r := 0
	for _, c := range t {
		r = dp[r][c-'a']
		if r == n {
			r = dp[0][c-'a']
			if r == n {
				Fprintln(out, -1)
				return
			}
			ans++
		}
		r++
	}

	Fprintln(out, ans+1)
}
func solve1(in io.Reader, out io.Writer) {
	var n, ans int
	var s, t string
	Fscan(in, &s, &t)
	n = len(s)
	const oz = 'z' + 1
	dp := make([][oz]int, n+1)
	//dp[n] = make([]int, oz)
	for i := 'a'; i <= 'z'; i++ {
		dp[n][i] = n
	}
	for i := n - 1; i >= 0; i-- {
		//dp[i] = make([]int, 26)
		for j := 'a'; j <= 'z'; j++ {
			dp[i][j] = dp[i+1][j]
		}
		dp[i][s[i]] = i
	}
	r := 0
	for _, c := range t {
		r = dp[r][c]
		if r == n {
			r = dp[0][c]
			if r == n {
				Fprintln(out, -1)
				return
			}
			ans++
		}
		r++
	}

	Fprintln(out, ans+1)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
