package hello

import (
	"os"
	"testing"
)

func TestGetHello(t *testing.T) {
	expect := os.Getenv("TEST_EXPECT")
	result := GetHello("テストさん")
	if result != expect {
		t.Error("\n実際：", result, "\n理想：", expect)
	}
}
