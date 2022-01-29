package dc

import (
	. "github.com/lxinyucn/go/c"
)

func D更新控件内容(ID1, ID2 int64, 内容 string, 是否unicode bool) string {
	var aa string
	if 是否unicode == true {
		aa = C字节集到十六进制([]byte(内容))
	} else {
		aa = C字节集到十六进制([]byte(C编码_utf8到gbk(内容)))
	}
	return C格式化文本("EEB110%04X%04X", ID1, ID2) + aa + "FFFCFFFF"
}
func D读取控件内容(ID1, ID2 int64, 内容 string, 是否unicode bool) string {
	return C格式化文本("EEB111%04X%04XFFFCFFFF", ID1, ID2)
}
func D更新控件数值(ID1, ID2, 数值 int64) string {
	return C格式化文本("EEB110%04X%04X%08XFFFCFFFF", ID1, ID2, 数值)
}
func D更新控件角度(ID1, ID2, 开始角度, 结束角度 int64) string {
	if 结束角度 < 0 {
		结束角度 = C到整数(360 + 结束角度)
	}
	return C格式化文本("EEB110%04X%04X%04X%04XFFFCFFFF", ID1, ID2, 开始角度, 结束角度)
}
func D读取控件数值(ID1, ID2 int64) string {
	return C格式化文本("EEB110%04X%04XFFFCFFFF", ID1, ID2)
}
func D发送自定义(内容 string, 是否unicode bool) string {
	var aa string
	if 是否unicode == true {
		aa = C字节集到十六进制([]byte(内容))
	} else {
		aa = C字节集到十六进制([]byte(C编码_utf8到gbk(内容)))
	}
	return "5A" + aa + "A55AA5A5"
}
func D发送自定义1(内容, 头, 尾巴 string, 是否unicode bool) string {
	var aa string
	if 是否unicode == true {
		aa = C字节集到十六进制([]byte(内容))
	} else {
		aa = C字节集到十六进制([]byte(C编码_utf8到gbk(内容)))
	}
	return 头 + aa + 尾巴
}
func D鸣叫(毫秒 int64) string {
	return C格式化文本("EE61%02XFFFCFFFF", 毫秒)
}
