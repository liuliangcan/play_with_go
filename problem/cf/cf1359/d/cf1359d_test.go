package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [d]")
    testCases := [][2]string{
        {`5
5 -2 10 -1 4`,`6`},
{`8
5 2 5 3 -30 -30 6 9`,`10`},
{`3
-10 6 -15`,`0`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/1359/problem/D
// https://codeforces.com/contest/1359/problem/D/submit?taskScreenName=cf1359d            
