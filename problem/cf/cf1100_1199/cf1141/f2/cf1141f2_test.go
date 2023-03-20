package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f2]")
	testCases := [][2]string{
		{`7
4 1 2 2 1 5 3`, `3
7 7
2 3
4 5`},
		{`11
-5 -4 -3 -2 -1 0 1 2 3 4 5`, `2
3 4
1 1`},
		{`4
1 1 1 1`, `4
4 4
1 1
2 2
3 3`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1141/F2
// https://codeforces.com/problemset/problem/1141/F2/submit?taskScreenName=cf1141f2
