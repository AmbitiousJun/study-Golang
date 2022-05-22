// 通过UDP监听的服务端
package main

import (
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
	// 监听1200端口
	service := ":1200"
	// 解析成udp地址
	addr, err := net.ResolveUDPAddr("udp4", service)
	check(err)
	// 开启监听
	l, err := net.ListenUDP("udp4", addr)
	check(err)
	for {
		cache := make([]byte, 128)
		_, rAddr, err := l.ReadFromUDP(cache)
		check(err)
		log.Println("监听到一个udp连接...")
		_, err = l.WriteToUDP(cache, rAddr)
		check(err)
	}
}
