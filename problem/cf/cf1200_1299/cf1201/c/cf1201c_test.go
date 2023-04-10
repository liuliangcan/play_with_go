package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`3 2
1 3 5`, `5`},
		{`5 5
1 2 1 1 1`, `3`},
		{`7 7
4 1 2 4 3 4 4`, `5`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1201/C
// https://codeforces.com/problemset/problem/1201/C/submit?taskScreenName=cf1201c
