package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{`11
ADD 3
SAVE 1
ADD 4
SAVE 2
LOAD 1
DELETE
DELETE
LOAD 2
SAVE 1
LOAD 3
LOAD 1`, `3 3 4 4 3 -1 -1 4 4 -1 4`},
		{`21
ADD 4
ADD 3
DELETE
ADD 10
LOAD 7
SAVE 5
SAVE 5
ADD 4
ADD 4
ADD 5
SAVE 5
ADD 2
DELETE
ADD 1
SAVE 5
ADD 7
ADD 8
DELETE
ADD 4
DELETE
LOAD 5`, `4 3 4 10 -1 -1 -1 4 4 5 5 2 5 1 1 7 8 7 4 7 1`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc273/tasks/abc273_e
// https://atcoder.jp/contests/abc273/tasks/abc273_e/submit?taskScreenName=abc273_e
