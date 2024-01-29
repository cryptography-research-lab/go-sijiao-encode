package sijiao_encoder

import (
	"github.com/cryptography-research-lab/go-sijiao-encode/data"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 一个四角编码的数字可能会被映射到多个中文，所以这是一个一对多的关系
// [四角编码的int数字, 映射到的中文字符的rune[]切片]
var siJiaoNumberToChineseMap = make(map[int][]rune)

// 一个中文只会被映射到一个四角编码的数字，但是这些数字可能会重复，也就是可能会出现同一个四角编码数字对应多个不同的中文的情况
// 所以这是一个一对一的关系
var chineseToSiJiaoNumberMap = make(map[rune]int)

// 初始化四角编码的字典映射
func init() {
	// 数据样例，有两列，第一列是中文汉字，第二列是中文对应的四角编码数字：
	// 稹 24981
	// 稺 27941
	// 稴 28937
	// 穃 23968
	// 穂 25937
	lines := strings.Split(data.Sijiaobianma2w, "\n")
	for _, line := range lines {
		split := strings.Split(strings.TrimSpace(line), " ")
		if len(split) != 2 {
			continue
		}

		// 期望第一列是一个汉字
		if utf8.RuneCountInString(split[0]) != 1 {
			continue
		}
		chinese := ([]rune(split[0]))[0]

		// 期望第二列是一个四角编码的数字
		siJiaoNumberString := split[1]
		siJiaoNumber, err := strconv.Atoi(siJiaoNumberString)
		if err != nil {
			// TODO 2024-1-19 21:36:46 此处是直接忽略比较好还是panic比较好呢？
			continue
		}

		siJiaoNumberToChineseMap[siJiaoNumber] = append(siJiaoNumberToChineseMap[siJiaoNumber], chinese)
		chineseToSiJiaoNumberMap[chinese] = siJiaoNumber
	}
}
