package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`3 3
5 4 6`, `5`},
		{`6 3
1 2 3 2 2 3`, `3`},
		{`6 3
1 2 4 1 2 3`, `3`},
		{`7 3
50 17 81 25 42 39 96`, `92`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1661/D
// https://codeforces.com/problemset/problem/1661/D/submit?taskScreenName=cf1661d
