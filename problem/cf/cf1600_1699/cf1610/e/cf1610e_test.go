package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{`4
3
1 2 3
5
1 4 4 5 6
6
7 8 197860736 212611869 360417095 837913434
8
6 10 56026534 405137099 550504063 784959015 802926648 967281024`, `0
1
2
3`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1610/E
// https://codeforces.com/problemset/problem/1610/E/submit?taskScreenName=cf1610e
