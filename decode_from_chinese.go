package sijiao_encoder

import (
	"sort"
)

// DecodeSiJiaoToChinese 根据中文查询其对应的四角编码
func DecodeSiJiaoToChinese(siJiaoNumber int) []rune {
	return siJiaoNumberToChineseMap[siJiaoNumber]
}

// TODO 2024-1-19 22:22:46 实现这个树型的方法
//// DecodeToChineseTrieFromSiJiaoSlice 解码为一颗中文的字典树
//func DecodeToChineseTrieFromSiJiaoSlice(siJiaoNumber []int) trie.Trie[rune] {
//	//
//	return siJiaoNumberToChineseMap[siJiaoNumber]
//}

// DecodeSiJiaoSliceToChineseRunes 根据中文查询其所有的可能对应的四角编码，注意，这个解码空间可能会比较大
func DecodeSiJiaoSliceToChineseRunes(siJiaoNumberSlice []int) [][]rune {
	resultSlice := make([][]rune, len(siJiaoNumberSlice))
	for index, sijiaoNumber := range siJiaoNumberSlice {
		chineseRunes := DecodeSiJiaoToChinese(sijiaoNumber)
		resultSlice[index] = chineseRunes
	}
	return resultSlice
}

// DecodeSiJiaoSliceToChineseString 对一连串四角编码进行解码，返回所有可能的中文组合
func DecodeSiJiaoSliceToChineseString(siJiaoNumberSlice []int) []string {

	chineseRunesSlice := DecodeSiJiaoSliceToChineseRunes(siJiaoNumberSlice)

	allMaybeChineseStringRunes := make([]string, 0)
	allMaybeChineseStringRunes = append(allMaybeChineseStringRunes, "")

	for _, allMaybeChineseRune := range chineseRunesSlice {
		newResultSlice := make([]string, 0)
		for _, chineseRune := range allMaybeChineseRune {
			for _, maybeChineseString := range allMaybeChineseStringRunes {
				newResultSlice = append(newResultSlice, maybeChineseString+string(chineseRune))
			}
		}
		allMaybeChineseStringRunes = newResultSlice
	}

	// 排序
	sort.Strings(allMaybeChineseStringRunes)

	return allMaybeChineseStringRunes
}
