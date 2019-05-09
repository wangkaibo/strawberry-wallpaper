package utils

import "strings"

func GetPlatformByUa(ua string) string {
	index := strings.Index(ua, "Macintosh")
	if index < 0 {
		return "windows"
	}

	return "mac"
}