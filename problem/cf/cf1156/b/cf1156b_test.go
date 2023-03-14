package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{`4
abcd
gg
codeforces
abaca`, `cadb
gg
codfoerces
No answer`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// http://codeforces.com/problemset/problem/1156/B
// http://codeforces.com/problemset/problem/1156/B/submit?taskScreenName=cf1156b
