package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`5
-1 1 -1 -1 1`, `Yes
-3 -1 4 5 7`},
		{`1
-1`, `Yes
0`},
		{`2
1 -1`, `No`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/arc153/tasks/arc153_c
// https://atcoder.jp/contests/arc153/tasks/arc153_c/submit?taskScreenName=arc153_c
