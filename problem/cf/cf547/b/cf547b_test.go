package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [b]")
    testCases := [][2]string{
        {`10
1 2 3 4 5 4 3 2 1 6`,`6 4 4 3 3 2 2 1 1 1`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/547/B
// https://codeforces.com/problemset/problem/547/B/submit?taskScreenName=cf547b            
