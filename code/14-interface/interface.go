// 接口的使用
package main

import "fmt"

// 定义一个手机接口
type Phone interface {
	// 打电话给某某人
	call(name string)
}

// 诺基亚结构体
type NokiaPhone struct{}

// 接口实现
func (phone NokiaPhone) call(name string) {
	fmt.Printf("使用诺基亚手机打电话给%s\n", name)
}

// 苹果手机结构体
type IPhone struct{}

// 接口实现
func (phone IPhone) call(name string) {
	fmt.Printf("使用苹果手机打电话给%s\n", name)
}

func main() {
	var phone Phone
	// 使用诺基亚实现
	phone = new(NokiaPhone)
	phone.call("张三")
	// 使用苹果手机实现
	phone = new(IPhone)
	phone.call("李四")
}
