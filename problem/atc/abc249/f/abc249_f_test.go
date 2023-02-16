package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [f]")
    testCases := [][2]string{
        {`5 1
2 4
2 -3
1 2
2 1
2 -3`,`3`},
{`1 0
2 -1000000000`,`-1000000000`},
{`10 3
2 3
2 -1
1 4
2 -1
2 5
2 -9
2 2
1 -6
2 5
2 -3`,`15`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc249/tasks/abc249_f
// https://atcoder.jp/contests/abc249/tasks/abc249_f/submit?taskScreenName=abc249_f            
