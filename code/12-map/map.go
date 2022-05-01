// map的使用
package main

import "fmt"

func main() {
	// 1 创建一个国家名称map
	countryMap := make(map[string]string)
	countryMap["France"] = "巴黎"
	countryMap["Italy"] = "意大利"
	countryMap["Japan"] = "东京"
	// 2 输出国家首都名
	for key := range countryMap {
		fmt.Printf("%s的首都是: %s\n", key, countryMap[key])
	}
	// 3 判断一个国家是否存在于map中
	capital, exist := countryMap["American"]
	fmt.Println(capital)
	fmt.Println(exist)
	if exist {
		fmt.Printf("American的首都是: %s\n", capital)
	} else {
		fmt.Println("American不在map集合中")
	}
}
