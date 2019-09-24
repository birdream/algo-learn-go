package learn

// 在一个字符串中找出没有重复字幕的最长串
// 任意ascii码
// 大小写敏感
func methodLSWRC(str string) string {
	l := 0
	r := -1
	length := len(str)
	keyMap := make(map[byte]int)
	var res string

	for l < length-1 {
		// fmt.Println("--------")
		// fmt.Println(keyMap[str[r]])
		// fmt.Println("--------")
		if r+1 < length-1 && keyMap[str[r+1]] == 0 {
			r++
			keyMap[str[r]]++
		} else {
			keyMap[str[l]]--
			l++
		}

		if len(res) < r-l+1 {
			res = str[l : r+1]
		}
	}

	return res
}
