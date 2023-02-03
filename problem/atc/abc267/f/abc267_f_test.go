package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{`5
1 2
2 3
3 4
3 5
3
2 2
5 3
3 3`, `4
1
-1`},
		{`10
1 2
2 3
3 5
2 8
3 4
4 6
4 9
5 7
9 10
5
1 1
2 2
3 3
4 4
5 5`, `2
4
10
-1
-1`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc267/tasks/abc267_f
// https://atcoder.jp/contests/abc267/tasks/abc267_f/submit?taskScreenName=abc267_f
