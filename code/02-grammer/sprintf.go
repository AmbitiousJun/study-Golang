package main

import (
	"fmt"
)

func main() {
	// %d表示整形数字，%s表示字符串
	var code = 123
	var enddate = "2022/4/30"
	var target = fmt.Sprintf("code=%d&endDate=%s", code, enddate)
	fmt.Println(target)
}