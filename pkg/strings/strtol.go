package strings

import (
	"strconv"
	"strings"
)

func Strtol(str string, base int) (int32, string) {
	str = strings.TrimLeft(str, " \t")
	i, err := strconv.ParseInt(str, base, 32)
	if err == nil {
		return int32(i), ""
	}
	validChars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	if (base == 0 || base == 16) && (str[0] == '0' && (str[1] == 'x' || str[1] == 'X')) {
		str = str[2:]
		base = 16
	}
	if base == 0 {
		if str[0] == '0' {
			base = 8
		} else {
			base = 10
		}
	}
	validChars = validChars[0:base]
	var res string
	for i, char := range str {
		if strings.ContainsAny(validChars, string(char)) == false {
			res = str[i:]
			str = str[0:i]
			break
		}
	}
	i, err = strconv.ParseInt(str, base, 32)
	if neg {
		i *= -1
	}
	return int32(i), res
}
