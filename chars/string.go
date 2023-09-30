package chars

import (
	`strings`
)

//
// Snake 驼峰转为串(蛇形)
//  @Description: XxYy to xx_yy , XxYY to xx_y_y
//  @param s 需要转换的字符串
//  @param separators 分隔驼峰字符串
//  @return string
//
func Snake(s string, separators ...byte) string {
	data := make([]byte, 0)
	j := false
	num := len(s)
	var separator byte = '_'
	if len(separators) > 0 {
		separator = separators[0]
	}
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, separator)
		}
		if d != separator {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

//
// Camel 蛇形转驼峰
//  @Description: xx_yy to XxYx  xx_y_y to XxYY
//  @param s 需要转换的字符串
//  @param separators 分隔驼峰字符串
//  @return string
//
func Camel(s string, separators ...byte) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	var separator byte = '_'
	if len(separators) > 0 {
		separator = separators[0]
	}
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == separator && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
