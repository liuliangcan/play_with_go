package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{`4
0 1 2 1
3 2 1 1`, `1 0 0 2`},
		{`7
2 5 1 5 3 4 3
2 4 3 5 6 5 1`, `0 0 0 1 0 2 4`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1157/E
// https://codeforces.com/problemset/problem/1157/E/submit?taskScreenName=cf1157e
