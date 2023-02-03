package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{`5
5 8 8 6 6`, `3`},
		{`7
12 8 11 3 3 13 2`, `3`},
		{`10
8 6 3 9 5 4 7 2 1 10`, `1`},
		{`14
630551244 683685976 249199599 863395255 667330388 617766025 564631293 614195656 944865979 277535591 390222868 527065404 136842536 971731491`, `8`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由XXX自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/arc119/tasks/arc119_c
// https://atcoder.jp/contests/arc119/tasks/arc119_c/submit?taskScreenName=arc119_c
