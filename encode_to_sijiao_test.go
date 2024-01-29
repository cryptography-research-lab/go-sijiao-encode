package sijiao_encoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeChineseRuneToSiJiao(t *testing.T) {
	sijiao := EncodeChineseRuneToSiJiao('陈')
	assert.Equal(t, 74294, sijiao)
}

func TestEncodeChineseRunesToSiJiao(t *testing.T) {
	sijiaoSlice := EncodeChineseRunesToSiJiao([]rune{'陈', '二'})
	assert.Equal(t, []int{74294, 10100}, sijiaoSlice)
}

func TestEncodeChineseStringToSiJiao(t *testing.T) {
	sijiaoSlice := EncodeChineseStringToSiJiao("陈二")
	assert.Equal(t, []int{74294, 10100}, sijiaoSlice)
}
