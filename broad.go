package websocket_manager

// 广播数据的结构体

type BroadCastMessageData struct {
	Message     []byte
	IsBroadCast bool
	ClientIDs   []string
}

// ping and pong
type PingPongMsg struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Alias   string `json:"alias"` // 别名
}
