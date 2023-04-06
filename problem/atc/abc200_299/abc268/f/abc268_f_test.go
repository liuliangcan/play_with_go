package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{`3
1X3
59
XXX`, `71`},
		{`10
X63X395XX
X2XX3X22X
13
3716XXX6
45X
X6XX
9238
281X92
1XX4X4XX6
54X9X711X1`, `3010`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc268/tasks/abc268_f
// https://atcoder.jp/contests/abc268/tasks/abc268_f/submit?taskScreenName=abc268_f
