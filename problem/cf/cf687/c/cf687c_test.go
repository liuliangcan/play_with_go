package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`6 18
5 6 1 10 12 2`, `16
0 1 2 3 5 6 7 8 10 11 12 13 15 16 17 18`},
		{`3 50
25 25 50`, `3
0 25 50`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/687/C
// https://codeforces.com/problemset/problem/687/C/submit?taskScreenName=cf687c
