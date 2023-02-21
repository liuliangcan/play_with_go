package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [d]")
    testCases := [][2]string{
        {`3 2
1 3 1`,`5`},
{`3 6
3 3 3`,`12`},
{`5 6
4 2 3 1 3`,`15`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/1358/D
// https://codeforces.com/problemset/problem/1358/D/submit?taskScreenName=cf1358d            
