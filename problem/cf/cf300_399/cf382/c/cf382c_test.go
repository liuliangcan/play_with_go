package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`3
4 1 7`, `2
-2 10`},
		{`1
10`, `-1`},
		{`4
1 3 5 9`, `1
7`},
		{`4
4 3 4 5`, `0`},
		{`2
2 4`, `3
0 3 6`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/382/C
// https://codeforces.com/problemset/problem/382/C/submit?taskScreenName=cf382c
