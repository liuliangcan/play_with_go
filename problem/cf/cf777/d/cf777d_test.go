package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{`3
#book
#bigtown
#big`, `#b
#big
#big`},
		{`3
#book
#cool
#cold`, `#book
#co
#cold`},
		{`4
#car
#cart
#art
#at`, `#
#
#art
#at`},
		{`3
#apple
#apple
#fruit`, `#apple
#apple
#fruit`},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

// 本模板由https://github.com/liuliangcan/play_with_python/blob/main/tools/gen_code_tool/gen_template.py自动生成,在本目录下使用go test进行case测试
// https://codeforces.com/problemset/problem/777/D
// https://codeforces.com/problemset/problem/777/D/submit?taskScreenName=cf777d
