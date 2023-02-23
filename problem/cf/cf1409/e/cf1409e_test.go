package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [e]")
    testCases := [][2]string{
        {`4
7 1
1 5 2 3 1 5 4
1 3 6 7 2 5 4
1 1
1000000000
1000000000
5 10
10 7 5 15 8
20 199 192 219 1904
10 10
15 19 8 17 20 10 9 2 10 19
12 13 6 17 1 14 7 9 19 3`,`6
1
5
10`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1409/E
// https://codeforces.com/problemset/problem/1409/E/submit?taskScreenName=cf1409e            
