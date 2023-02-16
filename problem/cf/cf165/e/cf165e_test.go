package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [e]")
    testCases := [][2]string{
        {`2
90 36`,`36 90`},
{`4
3 6 3 6`,`-1 -1 -1 -1`},
{`5
10 6 9 8 2`,`-1 8 2 2 8`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/165/e
// https://codeforces.com/problemset/problem/165/e/submit?taskScreenName=cf165e            
