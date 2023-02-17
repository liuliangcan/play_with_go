package main

import (
    "github.com/EndlessCheng/codeforces-go/main/testutil"
    "testing"
)

func Test_run(t *testing.T) {
    t.Log("Current test is [f]")
    testCases := [][2]string{
        {`3
1 2 2
2 3 3
1 2 3`,`8
6
6`},
{`6
1 2 3
1 3 1
1 4 4
1 5 1
1 6 5
9 2 6 5 3 100`,`105
108
106
109
106
14`},
{`6
1 2 1000000000
2 3 1000000000
3 4 1000000000
4 5 1000000000
5 6 1000000000
1 2 3 4 5 6`,`5000000006
4000000006
3000000006
3000000001
4000000001
5000000001`},
    }
    testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://atcoder.jp/contests/abc222/tasks/abc222_f
// https://atcoder.jp/contests/abc222/tasks/abc222_f/submit?taskScreenName=abc222_f            
