package websocket_manager

import (
	"github.com/gorilla/websocket"
	"log"
)

/*
 * todo: 制作管理器所需要的客户端
 * author: jocker
 * date: 2023年7月10日11:25:58
 */

type Client struct {
	Id      string          `json:"id"`
	Conn    *websocket.Conn // 链接
	Message chan []byte     // 消息队列
	Alias   string          `json:"alias"` // 别名
}

func (c Client) Read(readHandler func(messageType int, p []byte)) {
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		readHandler(messageType, p)
	}
}

func (c Client) Write() {
	for true {
		select {
		case msg := <-c.Message:
			if err := c.Conn.WriteMessage(TextMessage, msg); err != nil {
				log.Println(err)
				return
			}
			log.Printf("send msg success, msg body is %s", string(msg))
		}
	}
}
