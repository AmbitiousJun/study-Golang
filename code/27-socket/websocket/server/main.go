// websocket server端
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("监听到客户端连接...")
	for {
		var msg string
		// 接收客户端的消息
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("fail to recive msg")
			break
		}
		log.Println("received msg: ", msg)
		if err := websocket.Message.Send(ws, "Server Response to msg:"+msg); err != nil {
			log.Println("fail to send response")
			break
		}
	}
}

func main() {
	// 设置处理函数
	http.Handle("/", websocket.Handler(Echo))
	// 监听1234端口
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal(err)
	}
}
