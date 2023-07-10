package websocket_manager

// 广播数据的结构体

type BroadCastMessageData struct {
	Message     []byte
	IsBroadCast bool
	ClientIDs   []string
}
