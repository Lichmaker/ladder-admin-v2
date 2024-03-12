package utils

import (
	"strings"

	"github.com/spf13/cast"
)

func ParseHostPort(str string) (string, uint64) {
	tmp := strings.Split(str, ":")
	switch len(tmp) {
	case 0:
		return "", 0
	case 1:
		return tmp[0], 0
	case 2:
		return tmp[0], cast.ToUint64(tmp[1])
	default:
		return "", 0
	}
}
