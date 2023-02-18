package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{`4 2
1 20 4 7
20 2 20 3
1 3 5
1 4 6`, `16`},
		{`3 1
1 1 1
10 10 10
1 2 100`, `3`},
		{`7 8
35 29 36 88 58 15 25
99 7 49 61 67 4 57
2 3 3
2 5 36
2 6 89
1 6 24
5 7 55
1 3 71
3 4 94
5 6 21`, `160`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc270/tasks/abc270_f
// https://atcoder.jp/contests/abc270/tasks/abc270_f/submit?taskScreenName=abc270_f
