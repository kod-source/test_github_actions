package hello

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetHello(t *testing.T) {
	err := godotenv.Load("./../.env")
	if err != nil {
		t.Error(err)
	}
	expect := os.Getenv("TEST_EXPECT")
	result := GetHello("テストさん")
	if result != expect {
		t.Error("\n実際：", result, "\n理想：", expect)
	}
}
