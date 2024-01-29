# 四角编码

TODO 2023-2-3 01:28:40 编写文档 

# 一、四角编码是什么？应用场景有哪些？

TODO 2024-1-29 21:59:38 

# 二、安装依赖

```bash
go get -u github.com/cryptography-research-lab/go-sijiao-encode 
```

# 三、API示例

```go
package main

import (
	"fmt"
	sijiao_encoder "github.com/cryptography-research-lab/go-sijiao-encode"
)

func main() {

	s := "陈二"

	// 对中文进行四角编码
	sijiaoSlice := sijiao_encoder.EncodeChineseStringToSiJiao(s)
	fmt.Println(sijiaoSlice)
	// Output:
	// [74294 10100]

	// 把四角编码解码为所有可能得中文字符串
	maybeChineseStrings := sijiao_encoder.DecodeSiJiaoSliceToChineseStrings(sijiaoSlice)
	fmt.Println(maybeChineseStrings)
	// Output:
	// [胨二 脎二 腜二 陈二]

	// 因为四角编码到中文的映射是一对多的，所以解码的结果可能会有很多，这个方法是把结果以字典树的方式返回
	trie, err := sijiao_encoder.DecodeToChineseTrieFromSiJiaoSlice(sijiaoSlice)
	if err != nil {
		panic(err)
	}
	fmt.Println(trie.ExportToDotLanguage())
	// Output:
	// digraph G1 {
	//    "0::" -> "1:腜:";
	//    "0::" -> "2:陈:";
	//    "0::" -> "3:胨:";
	//    "0::" -> "4:脎:";
	//    "1:腜:" -> "5:二:腜二";
	//    "2:陈:" -> "6:二:陈二";
	//    "3:胨:" -> "7:二:胨二";
	//    "4:脎:" -> "8:二:脎二";
	// }

}
```

# 四、TODO 

 从四角编码到中文的映射是一对多的，这是展开的字典树，看看是不是能够有什么办法，从这颗庞大的字典树中定位到正确的解？

![image-20240129220216393](README.assets/image-20240129220216393.png)

一个可能的思路：搞一个汉字矩阵表，存储这每个汉字的下一个汉字是啥的概率，比如汉字A的下一个汉字是A的概率，下一个汉字是B的概率，下一个汉字是C的概率，等等等等...然后计算字典树上所有叶子的概率，概率最大的那个就是最可能的答案。





