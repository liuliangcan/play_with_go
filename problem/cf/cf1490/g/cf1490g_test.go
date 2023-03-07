package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [g]")
	testCases := [][2]string{
		{`3
3 3
1 -3 4
1 5 2
2 2
-2 0
1 2
2 2
0 1
1 2`, `0 6 2 
-1 -1 
1 3`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1490/G
// https://codeforces.com/problemset/problem/1490/G/submit?taskScreenName=cf1490g
