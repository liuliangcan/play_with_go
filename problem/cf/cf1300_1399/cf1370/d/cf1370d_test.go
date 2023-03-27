package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`4 2
1 2 3 4`, `1`},
		{`4 3
1 2 3 4`, `2`},
		{`5 3
5 3 4 2 6`, `2`},
		{`6 4
5 3 50 2 4 5`, `3`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1370/D
// https://codeforces.com/problemset/problem/1370/D/submit?taskScreenName=cf1370d
