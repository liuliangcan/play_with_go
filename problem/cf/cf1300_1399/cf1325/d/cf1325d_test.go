package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`2 4`, `2
3 1`},
		{`1 3`, `3
1 1 1`},
		{`8 5`, `-1`},
		{`0 0`, `0`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1325/D
// https://codeforces.com/problemset/problem/1325/D/submit?taskScreenName=cf1325d
