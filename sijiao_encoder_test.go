package sijiao_encoder

import (
	"testing"
)

func TestEncode(t *testing.T) {
	s := "这是中文"
	encode := QuerySiJiaoForChineseString(s)
	t.Log(encode)
}
