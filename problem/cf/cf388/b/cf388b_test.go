package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{`2`, `4
NNYY
NNYY
YYNN
YYNN`},
		//{`9`,`8
		//NNYYYNNN
		//NNNNNYYY
		//YNNNNYYY
		//YNNNNYYY
		//YNNNNYYY
		//NYYYYNNN
		//NYYYYNNN
		//NYYYYNNN`},
		//{`1`,`2
		//NY
		//YN`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/388/B
// https://codeforces.com/problemset/problem/388/B/submit?taskScreenName=cf388b
