// 错误处理接口的使用
package main

import (
	"errors"
	"fmt"
	"math"
)

// 对一个数取根号
// 返回两个参数，一个是取根号后的结果，另一个是error接口类型
func Sqrt(n int) (float64, error) {
	if n < 0 {
		// 负数无法求根号
		return 0, errors.New("负数不能取根号")
	}
	return math.Sqrt(float64(n)), nil
}

// 除法异常结构体
type DevideError struct {
	devidee int
	devider int
}

// 除法异常实现error接口
func (de DevideError) Error() string {
	msg := `
		除法运算执行错误, 除数不能为0
		devidee: %d
		devider: 0
	`
	return fmt.Sprintf(msg, de.devidee)
}

// 除法函数
// func Devide(devidee, devider int) (float64, *DevideError) {
// 	if devider == 0 {
// 		return 0, &DevideError{devidee, devider}
// 	}
// 	return float64(devidee / devider), nil
// }

func Devide(devidee, devider int) (float64, error) {
	if devider == 0 {
		return 0, DevideError{devidee, devider}
	}
	return float64(devidee / devider), nil
}

func main() {
	// 测试1：传入正常的值求根号
	ans1, err1 := Sqrt(25)
	if err1 == nil {
		fmt.Printf("25取根号的结果为: %.3f\n", ans1)
	} else {
		fmt.Println(err1)
	}
	// 测试2：传入异常值
	ans2, err2 := Sqrt(-3)
	if err2 == nil {
		fmt.Printf("-3取根号的结果为: %.3f\n", ans2)
	} else {
		fmt.Println(err2)
	}

	fmt.Println("==================================")

	// 测试3：正常除法
	ans3, err3 := Devide(9, 3)
	if err3 == nil {
		fmt.Printf("9 / 3 的结果为: %.3f\n", ans3)
	} else {
		fmt.Println(err3)
	}
	// 测试4：异常除法
	ans4, err4 := Devide(9, 0)
	if err4 == nil {
		fmt.Printf("9 / 0 的结果为: %.3f\n", ans4)
	} else {
		fmt.Println(err4)
	}
}
