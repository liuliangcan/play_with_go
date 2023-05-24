package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`5
1 3 4 3 3`, `4 2 1 1 1`},
		{`8
1 5 7 8 1 7 6 1`, `8 4 3 2 1 1 1 1`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/786/problem/C
// https://codeforces.com/contest/786/problem/C/submit?taskScreenName=cf786c
