package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{`5 3 2`, `1 2
1 3
3 4
3 5`},
		{`8 5 2`, `-1`},
		{`8 4 2`, `4 8
5 7
2 3
8 1
2 1
5 6
1 5`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/639/B
// https://codeforces.com/problemset/problem/639/B/submit?taskScreenName=cf639b
