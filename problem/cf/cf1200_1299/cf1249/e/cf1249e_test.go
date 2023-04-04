package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{`10 2
7 6 18 6 16 18 1 17 17
6 9 3 10 9 1 10 1 5`, `0 7 13 18 24 35 36 37 40 45`},
		{`10 1
3 2 3 1 3 3 1 4 1
1 2 3 4 4 1 2 1 3`, `0 2 4 7 8 11 13 14 16 17`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1249/E
// https://codeforces.com/problemset/problem/1249/E/submit?taskScreenName=cf1249e
