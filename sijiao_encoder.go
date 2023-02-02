package sijiao_encoder

// QuerySiJiaoForChinese 根据四角编码查询对应的中文
func QuerySiJiaoForChinese(chineseCharacter rune) int {
	return chineseToSiJiaoNumberMap[chineseCharacter]
}

// QueryChineseForSiJiao 根据中文查询其对应的四角编码
func QueryChineseForSiJiao(siJiaoNumber int) []rune {
	return siJiaoNumberToChineseMap[siJiaoNumber]
}

// QuerySiJiaoForChineseString 对中文进行四角编码
func QuerySiJiaoForChineseString(chinese string) []int {
	chineseRuneSlice := []rune(chinese)
	resultSlice := make([]int, len(chineseRuneSlice))
	for index, character := range chineseRuneSlice {
		resultSlice[index] = QuerySiJiaoForChinese(character)
	}
	return resultSlice
}
