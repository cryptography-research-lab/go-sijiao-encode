package sijiao_encoder

// EncodeChineseRuneToSiJiao 根据四角编码查询对应的中文
func EncodeChineseRuneToSiJiao(chineseCharacter rune) int {
	return chineseToSiJiaoNumberMap[chineseCharacter]
}

// EncodeChineseStringToSiJiao 对一串中文进行四角编码
func EncodeChineseStringToSiJiao(chinese string) []int {
	chineseRuneSlice := []rune(chinese)
	return EncodeChineseRunesToSiJiao(chineseRuneSlice)
}

// EncodeChineseRunesToSiJiao 把中文字符串转为四角编码切片
func EncodeChineseRunesToSiJiao(chineseRuneSlice []rune) []int {
	resultSlice := make([]int, len(chineseRuneSlice))
	for index, character := range chineseRuneSlice {
		resultSlice[index] = EncodeChineseRuneToSiJiao(character)
	}
	return resultSlice
}
