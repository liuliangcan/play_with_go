package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [e]")
    testCases := [][2]string{
        {`7 24 21 23
16 17 14 20 20 11 22`,`3`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/contest/1324/problem/E
// https://codeforces.com/contest/1324/problem/E/submit?taskScreenName=cf1324e            
