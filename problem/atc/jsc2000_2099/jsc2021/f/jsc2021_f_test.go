package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{`2 2 4
1 1 10
2 1 20
2 2 5
1 1 30`, `20
50
55
85`},
		{`3 3 5
1 3 10
2 1 7
1 3 5
2 2 10
2 1 1`, `30
44
31
56
42`},
		{`200000 200000 4
2 112219 100000000
1 73821 100000000
2 26402 100000000
1 73821 100000000`, `20000000000000
39999900000000
59999800000000
59999800000000`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/jsc2021/tasks/jsc2021_f
// https://atcoder.jp/contests/jsc2021/tasks/jsc2021_f/submit?taskScreenName=jsc2021_f
