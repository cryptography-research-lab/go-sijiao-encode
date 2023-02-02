package sijiao_encoder

import (
	"github.com/cryptography-research-lab/go-sijiao-encode/data"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 一个四角编码的数字可能会被映射到多个中文
var siJiaoNumberToChineseMap = make(map[int][]rune)

// 一个中文只会被映射到一个四角编码的数字，但是这些数字可能会重复，也就是可能会出现同一个四角编码数字对应多个不同的中文的情况
var chineseToSiJiaoNumberMap = make(map[rune]int)

// 初始化四角编码的字典映射
func init() {
	lines := strings.Split(data.Sijiaobianma2w, "\n")
	for _, line := range lines {
		split := strings.Split(strings.TrimSpace(line), " ")
		if len(split) != 2 {
			continue
		}

		if utf8.RuneCountInString(split[0]) != 1 {
			continue
		}
		chinese := ([]rune(split[0]))[0]

		siJiaoNumberString := split[1]
		siJiaoNumber, err := strconv.Atoi(siJiaoNumberString)
		if err != nil {
			continue
		}

		siJiaoNumberToChineseMap[siJiaoNumber] = append(siJiaoNumberToChineseMap[siJiaoNumber], chinese)
		chineseToSiJiaoNumberMap[chinese] = siJiaoNumber
	}
}
