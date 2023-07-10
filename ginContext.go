package websocket_manager

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	defaultManager *Manager
	defaultPool    *Pool
)

func init() {
	mUuidStr := uuid.New().String()
	defaultManager = InitializationManager(mUuidStr)

	defaultPool = InitializationPool()
	// 将默认房间注册进连接池
	defaultPool.Register <- defaultManager
}

// 初始化 gin 的上下文
// 欠缺ping pond的回复

func WsGinContext(ctx *gin.Context) {
	upGrande := websocket.Upgrader{
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		//设置请求协议
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	// 获取某个参数 类似user_id
	tempStr := ctx.Query("user_id")

	//创建连接
	conn, err := upGrande.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("websocket connect error: %s", tempStr)
		return
	}

	//生成唯一标识client_id
	uuidStr := uuid.New().String()

	client := &Client{
		Id:      uuidStr,
		Conn:    conn,
		Message: make(chan []byte, 2048),
	}

	// 如果没有申明房间号，就进入默认房间
	defaultManager.Register <- client

	//起协程，实时接收和回复数据
	go client.Read(func(messageType int, bt []byte) {
		log.Printf("received message, msg type is %d, body is : %s", messageType, string(bt))
	})
	go client.Write()
}
