package learn

import (
	"regexp"
	"strings"
)

// 判断是否回文串 忽略大小写 只看数字和字母
func validPalindrome(str string) bool {
	// str = strings.ToLower(str)
	i := 0
	j := len(str) - 1
	pattern := "[a-z0-9A-Z]"

	for i < j {
		if ok, _ := regexp.MatchString(pattern, string(str[i])); !ok {
			i++
			continue
		}

		if ok, _ := regexp.Match(pattern, []byte{str[j]}); !ok {
			j--
			continue
		}

		if strings.EqualFold(string(str[i]), string(str[j])) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
