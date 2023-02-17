package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [a]")
    testCases := [][2]string{
        {`3
BBA`,`Yes`},
{`4
ABAB`,`No`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/arc145/tasks/arc145_a
// https://atcoder.jp/contests/arc145/tasks/arc145_a/submit?taskScreenName=arc145_a            
