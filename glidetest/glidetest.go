package main

import (
	"fmt"
	"github.com/sh3lld00m/go-sample-package/samplepackage"
)

func main() {
	fmt.Println("Version: ", samplepackage.GetBranch())
}
