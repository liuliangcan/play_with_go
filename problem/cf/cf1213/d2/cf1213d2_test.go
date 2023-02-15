package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [d2]")
    testCases := [][2]string{
        {`5 3
1 2 2 4 5`,`1`},
{`5 3
1 2 3 4 5`,`2`},
{`5 3
1 2 3 3 3`,`0`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1213/D2
// https://codeforces.com/problemset/problem/1213/D2/submit?taskScreenName=cf1213d2            
