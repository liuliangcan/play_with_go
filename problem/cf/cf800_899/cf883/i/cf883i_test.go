package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [i]")
	testCases := [][2]string{
		{`5 2
50 110 130 40 120`, `20`},
		{`4 1
2 3 4 1`, `0`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/883/I
// https://codeforces.com/problemset/problem/883/I/submit?taskScreenName=cf883i
