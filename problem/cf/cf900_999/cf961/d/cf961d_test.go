package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`5
0 0
0 1
1 1
1 -1
2 2`, `YES`},
		{`5
0 0
1 0
2 1
1 1
2 3`, `NO`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/961/problem/D
// https://codeforces.com/contest/961/problem/D/submit?taskScreenName=cf961d
