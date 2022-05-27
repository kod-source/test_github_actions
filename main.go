package main

import (
	"fmt"

	"github.com/kod-source/test_github_actions/hello"
)

func main() {
	s := hello.GetHello("テスト君")
	fmt.Println(s)
}
