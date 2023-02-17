package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [e]")
    testCases := [][2]string{
        {`3 3
1 3 0`,`0 1 0`},
{`2 1000000000000
1000000000000 1000000000000`,`500000000000 500000000000`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc270/tasks/abc270_e
// https://atcoder.jp/contests/abc270/tasks/abc270_e/submit?taskScreenName=abc270_e            
