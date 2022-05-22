// 通过UDP连接的客户端
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	// 连接到本地的1200端口
	service := ":1200"
	// 解析成udp地址
	addr, err := net.ResolveUDPAddr("udp4", service)
	check(err)
	// 取得udp连接
	conn, err := net.DialUDP("udp4", nil, addr)
	check(err)
	_, err = conn.Write([]byte("Hello World!!!!!"))
	check(err)
	cache := make([]byte, 128)
	_, err = conn.Read(cache)
	check(err)
	fmt.Println(string(cache))
}
