package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`6
3 5 4 7 10 12`, `3
7 3 5`},
		{`5
-1 2 5 8 11`, `1
8`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/988/problem/D
// https://codeforces.com/contest/988/problem/D/submit?taskScreenName=cf988d
