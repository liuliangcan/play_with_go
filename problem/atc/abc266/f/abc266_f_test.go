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
1 3
1 4
2 5
3
1 2
1 4
1 5`,`No
Yes
No`},
{`10
3 5
5 7
4 8
2 9
1 2
7 9
1 6
4 10
2 5
2 10
10
1 8
6 9
8 10
6 8
3 10
3 9
1 10
5 8
1 10
7 8`,`Yes
No
Yes
Yes
No
No
Yes
No
Yes
No`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc266/tasks/abc266_f
// https://atcoder.jp/contests/abc266/tasks/abc266_f/submit?taskScreenName=abc266_f            
