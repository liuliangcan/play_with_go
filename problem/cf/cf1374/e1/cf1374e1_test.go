package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [e1]")
    testCases := [][2]string{
        {`8 4
7 1 1
2 1 1
4 0 1
8 1 1
1 0 1
1 1 1
1 0 1
3 0 0`,`18`},
{`5 2
6 0 0
9 0 0
1 0 1
2 1 1
5 1 0`,`8`},
{`5 3
3 0 0
2 1 0
3 1 0
5 0 1
3 0 1`,`-1`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/1374/problem/E1
// https://codeforces.com/contest/1374/problem/E1/submit?taskScreenName=cf1374e1            
