package utils

import "strings"

// MakeUrl 生成字符串
func MakeUrl(baseUrl, token, message string) string {
	return JoinString(baseUrl, "/", token, "/", message)
}

// JoinString 拼接字符串
func JoinString(multiStr ...string) string {
	var sb strings.Builder
	for _, str := range multiStr {
		sb.WriteString(str)
	}
	return sb.String()
}
