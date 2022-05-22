// 通过TCP 连接的客户端
package main

import (
	"fmt"
	"net"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	service := ":1200"
	// 转化为tcp地址
	addr, err := net.ResolveTCPAddr("tcp4", service)
	check(err)
	// 建立连接
	conn, err := net.DialTCP("tcp4", nil, addr)
	check(err)
	defer conn.Close()
	fmt.Println("成功连接服务器...")

	fmt.Println("发送timestamp请求...")
	conn.Write([]byte("timestamp"))
	cache := make([]byte, 128)
	_, err = conn.Read(cache)
	check(err)
	fmt.Printf("timestamp: %s\n", string(cache))

	fmt.Println("发送normal请求...")
	conn.Write([]byte("normal"))
	cache = make([]byte, 128)
	_, err = conn.Read(cache)
	check(err)
	fmt.Printf("normal: %s\n", string(cache))
}
