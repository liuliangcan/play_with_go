package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [a]")
    testCases := [][2]string{
        {`4
3
1 2 1
5
11 7 9 6 8
5
1 3 1 3 1
4
5 2 1 10`,`YES
YES
NO
YES`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/1442/problem/A
// https://codeforces.com/contest/1442/problem/A/submit?taskScreenName=cf1442a            
