package sijiao_encoder

import (
	"github.com/golang-infrastructure/go-trie"
	"sort"
)

// DecodeSiJiaoToChinese 根据中文查询其对应的四角编码
func DecodeSiJiaoToChinese(siJiaoNumber int) []rune {
	return siJiaoNumberToChineseMap[siJiaoNumber]
}

// DecodeSiJiaoSliceToChineseRunes 根据中文查询其所有的可能对应的四角编码，注意，这个解码空间可能会比较大
func DecodeSiJiaoSliceToChineseRunes(siJiaoNumberSlice []int) [][]rune {
	resultSlice := make([][]rune, len(siJiaoNumberSlice))
	for index, sijiaoNumber := range siJiaoNumberSlice {
		chineseRunes := DecodeSiJiaoToChinese(sijiaoNumber)
		resultSlice[index] = chineseRunes
	}
	return resultSlice
}

// DecodeSiJiaoSliceToChineseStrings 对一连串四角编码进行解码，返回所有可能的中文组合
func DecodeSiJiaoSliceToChineseStrings(siJiaoNumberSlice []int) []string {

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

// DecodeToChineseTrieFromSiJiaoSlice 解码为一颗中文的字典树
func DecodeToChineseTrieFromSiJiaoSlice(siJiaoSlice []int) (*trie.Trie[string], error) {
	trie := trie.New[string](func(s string) ([]string, error) {
		runes := []rune(s)
		resultSlice := make([]string, len(runes))
		for i := 0; i < len(runes); i++ {
			resultSlice[i] = string(runes[i])
		}
		return resultSlice, nil
	})
	strings := DecodeSiJiaoSliceToChineseStrings(siJiaoSlice)
	for _, s := range strings {
		if err := trie.Add(s, s); err != nil {
			return nil, err
		}
	}
	return trie, nil
}
