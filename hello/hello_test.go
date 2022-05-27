package hello

import "testing"

func TestGetHello(t *testing.T) {
	result := GetHello("テストさん")
	expect := "こんにちは、テストさん！"
	if result != expect {
		t.Error("\n実際：", result, "\n理想：", expect)
	}
}
