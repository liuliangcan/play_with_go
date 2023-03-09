package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`3 2 2 2`, `AC
AC`},
		{`3 3 2 2`, `Happy new year!`},
		//{`3 0 2 2`,`AA
		//AA`},
		{`4 3 2 1`, `Happy new year!`},
		{`4 2 2 1`, `Happy new year!`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/379/D
// https://codeforces.com/problemset/problem/379/D/submit?taskScreenName=cf379d
