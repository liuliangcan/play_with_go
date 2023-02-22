package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [c]")
    testCases := [][2]string{
        {`3 2
abc
xyc`,`ayd`},
{`1 0
c
b`,`-1`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/584/C
// https://codeforces.com/problemset/problem/584/C/submit?taskScreenName=cf584c            
