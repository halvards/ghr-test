package main

import (
	"fmt"

	"github.com/halvards/ghr-test/cmd/version"
)

func main() {
	fmt.Printf("hello world %s\n", version.Get())
}
