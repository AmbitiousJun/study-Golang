// 通过TCP监听的服务端
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// 将当前系统时间响应给客户端
// 客户端发送"timestamp"，返回时间戳
// 客户端发送"normal"，返回正常的字符串时间
func handleClient(conn net.Conn) {
	// 设置接收消息的超时时间为两分钟
	conn.SetReadDeadline(time.Now().Add(time.Minute * 2))
	defer conn.Close()
	// 存放客户端发送过来的数据
	cache := make([]byte, 128)
	// 持续监听
	for {
		readLen, err := conn.Read(cache)
		check(err)
		if readLen == 0 {
			// 客户端主动关闭了连接
			break
		}
		msg := string(cache[:readLen])
		fmt.Printf("msg: %v\n", msg)
		var timeStr string
		if msg == "timestamp" {
			timeStr = strconv.FormatInt(time.Now().Unix(), 10)
		} else {
			timeStr = time.Now().String()
		}
		conn.Write([]byte(timeStr))
		// 清空缓存
		cache = make([]byte, 128)
	}
}

func main() {
	// 监听本地的1200端口
	service := ":1200"
	// 解析地址信息
	addr, err := net.ResolveTCPAddr("tcp4", service)
	check(err)
	// 开启服务监听
	l, err := net.ListenTCP("tcp4", addr)
	check(err)
	fmt.Println("开始监听...")
	// 持续监听
	for {
		// 接收客户端的连接请求
		conn, err := l.Accept()
		check(err)
		fmt.Println("监听到一个连接...")
		go handleClient(conn)
	}
}
