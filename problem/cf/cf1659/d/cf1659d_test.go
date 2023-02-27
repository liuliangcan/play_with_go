package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`5
4
2 4 2 4
7
0 3 4 2 3 2 7
3
0 0 0
4
0 0 0 4
3
1 2 3`, `1 1 0 1 
0 1 1 0 0 0 1 
0 0 0 
0 0 0 1 
1 0 1`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1659/D
// https://codeforces.com/problemset/problem/1659/D/submit?taskScreenName=cf1659d
