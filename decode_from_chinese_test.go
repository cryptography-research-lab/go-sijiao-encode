package sijiao_encoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeSiJiaoToChinese(t *testing.T) {
	chinese := DecodeSiJiaoToChinese(74294)
	assert.Equal(t, []int32{38472, 33000, 33038, 33116}, chinese)
}

func TestDecodeSiJiaoSliceToChineseRunes(t *testing.T) {
	sijiaoSlice := EncodeChineseStringToSiJiao("编码的快乐是纯粹的快乐")
	chinese := DecodeSiJiaoSliceToChineseRunes(sijiaoSlice)
	//marshal, _ := json.Marshal(chinese)
	//fmt.Println(string(marshal))
	assert.Equal(t, [][]int32{{32534, 40138}, {37045, 30720, 30721, 37089, 37082, 37137, 30830, 30823, 30882, 30918, 30911, 37265, 30966, 40280, 37187, 40349}, {21477, 26092, 30008, 21260, 30340, 35335, 21264, 21266, 30359, 32763}, {24555, 24591, 24610}, {20048, 39652, 27267, 27389}, {36275, 26159, 30064}, {32424, 32431, 40064}, {31929}, {21477, 26092, 30008, 21260, 30340, 35335, 21264, 21266, 30359, 32763}, {24555, 24591, 24610}, {20048, 39652, 27267, 27389}}, chinese)
}

func TestDecodeSiJiaoSliceToChineseString(t *testing.T) {
	sijiaoSlice := EncodeChineseStringToSiJiao("加密算法")
	chinese := DecodeSiJiaoSliceToChineseString(sijiaoSlice)
	//marshal, _ := json.Marshal(chinese)
	//s := strings.ReplaceAll(string(marshal), "[", "{")
	//s = strings.ReplaceAll(s, "]", "}")
	//fmt.Println(s)
	assert.Equal(t, []string{"加宻算汯", "加宻算法", "加宻算浓", "加宻算溒", "加宻算濍", "加宻算濛", "加宻算瀡", "加密算汯", "加密算法", "加密算浓", "加密算溒", "加密算濍", "加密算濛", "加密算瀡", "加寚算汯", "加寚算法", "加寚算浓", "加寚算溒", "加寚算濍", "加寚算濛", "加寚算瀡", "加崈算汯", "加崈算法", "加崈算浓", "加崈算溒", "加崈算濍", "加崈算濛", "加崈算瀡", "加窋算汯", "加窋算法", "加窋算浓", "加窋算溒", "加窋算濍", "加窋算濛", "加窋算瀡", "加窑算汯", "加窑算法", "加窑算浓", "加窑算溒", "加窑算濍", "加窑算濛", "加窑算瀡", "加窰算汯", "加窰算法", "加窰算浓", "加窰算溒", "加窰算濍", "加窰算濛", "加窰算瀡"}, chinese)
}
