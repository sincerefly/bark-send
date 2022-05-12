package common

import (
	"bark-send/utils"
	"strings"
	"time"
)

var (
	Build string
	Branch string
	Commit string
	Platform string
)

// BuildVersion 版本信息
func BuildVersion() string {
	var buf strings.Builder

	if Build != "" {
		buf.WriteString("build: ")
		buf.WriteString(Build)
	}
	if Branch != "" {
		buf.WriteByte('\n')
		buf.WriteString("branch: ")
		buf.WriteString(Branch)
	}
	if Commit != "" {
		buf.WriteByte('\n')
		buf.WriteString("commit: ")
		buf.WriteString(Commit)
	}
	if Platform != "" {
		buf.WriteByte('\n')
		buf.WriteString("platform: ")
		buf.WriteString(Platform)
	}
	buf.WriteByte('\n')
	buf.WriteString("ServerTime: ")
	buf.WriteString(utils.CurrentTime().Format(time.RFC3339))
	return buf.String()
}
