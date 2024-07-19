package policy

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

func TestNew(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			jsonStr = `{
				"Black": {
					"BlackOpen": true,
					"BlackDay": 30,
					"BlackCount": 3,
					"BlackCycle": 365
				},
				"Hello": {
					"World": "Hello World"
				}
			}`
		)
		_, err := New(jsonStr)
		t.AssertNil(err)
		pf, err := Format(jsonStr)
		t.AssertNil(err)

		// 判断是否包含 PolicyTimeConf
		if !gstr.Contains(pf, "PolicyTimeConf") {
			t.Fatal("json 转码失败")
		}
		// 判断是否不包含 Hello
		if gstr.Contains(pf, "Hello") {
			t.Fatal("json 转码失败")
		}
	})
}
