package main

import (
	"bark-send/common"
	"bark-send/prepare"
	"fmt"
	"github.com/spf13/pflag"
)

func main() {

	versionInfo := pflag.BoolP("version", "v", false, "show version info.")
	pflag.Parse()
	if versionInfo != nil && *versionInfo {
		fmt.Println(common.BuildVersion())
		return
	}

	prepare.Server()
}
