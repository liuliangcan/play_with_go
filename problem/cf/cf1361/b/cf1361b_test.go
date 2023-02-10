package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [b]")
    testCases := [][2]string{
        {`4
5 2
2 3 4 4 3
3 1
2 10 1000
4 5
0 1 1 100
1 8
89`,`4
1
146981438
747093407`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1361/B
// https://codeforces.com/problemset/problem/1361/B/submit?taskScreenName=cf1361b            
