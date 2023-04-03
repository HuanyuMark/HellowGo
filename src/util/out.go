package util

import "fmt"

func GetInt() int {
	return 1
}

func testIfprotected() {
	fmt.Printf("out.go call test: %d", GetInt())
}
