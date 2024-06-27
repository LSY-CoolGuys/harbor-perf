package main

import (
	"fmt"
	"github.com/magefile/mage/mage"
)

// This file allows someone to run mage commands without mage installed
// by running `go run mage.go TARGET`.
// See https://magefile.org/zeroinstall/
func main() {
	if flags := mage.Main(); flags != 0 {
		fmt.Printf("测试发生错误，退出码%d")
	}
}
