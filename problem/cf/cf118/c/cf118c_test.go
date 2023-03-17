package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`6 5
898196`, `4
888188`},
		{`3 2
533`, `0
533`},
		{`10 6
0001112223`, `3
0000002223`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/118/C
// https://codeforces.com/problemset/problem/118/C/submit?taskScreenName=cf118c
