package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [b]")
    testCases := [][2]string{
        {`4
1 2
1 3
4 2
2 3`,`4`},
{`2
111 111
111 111`,`1`},
{`12
5 2
5 6
1 2
9 7
2 7
5 5
4 2
6 7
2 2
7 8
9 7
1 8`,`8`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/arc111/tasks/arc111_b
// https://atcoder.jp/contests/arc111/tasks/arc111_b/submit?taskScreenName=arc111_b            
