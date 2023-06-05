package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{`3 2`, `5`},
		{`6 4`, `39`},
		{`2 1`, `2`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/414/problem/B
// https://codeforces.com/contest/414/problem/B/submit?taskScreenName=cf414b
