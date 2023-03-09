package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`4
1 8
2 3
4 7
5 6`, `3
0
1
0`},
		{`3
		3 4
		1 5
		2 6`, `0
		1
		1`},
		{`10
-3 -1
-10 4
0 6
-4 -2
1 3
2 9
5 10
-7 -6
-8 8
-9 7`, `0
4
1
0
0
0	
0
0
5
5`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/652/D
// https://codeforces.com/problemset/problem/652/D/submit?taskScreenName=cf652d
