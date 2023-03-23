package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{`5
4 8
12340219
20215601
56782022
12300678
12345678
2 3
134
126
123
1 4
1210
1221
4 3
251
064
859
957
054
4 7
7968636
9486033
4614224
5454197
9482268`, `3
1 4 1
5 6 2
3 4 3
-1
2
1 2 1
2 3 1
-1
3
1 3 2
5 6 3
3 4 1`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1624/E
// https://codeforces.com/problemset/problem/1624/E/submit?taskScreenName=cf1624e
