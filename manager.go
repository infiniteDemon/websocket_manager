package websocket_manager

import (
	"log"
	"sync"
)

/*
 * date: 2023年7月10日11:15:35
 * author: joker
 * todo: 编写一个websocket客户端管理器
 */

// websocket clients manager

type Manager struct {
	Id                   string `json:"id"` // manager id
	Group                map[string]*Client
	Lock                 sync.Mutex
	Register, UnRegister chan *Client
	BroadCastMessage     chan *BroadCastMessageData
	Close                chan int
	clientCount          uint //分组及客户端数量
}

func (m Manager) Run() {
	for {
		select {
		case num := <-m.Close:
			log.Printf("pool close signal = %d", num)
			break
		case manager := <-m.Register:
			//注册客户端
			m.Lock.Lock()
			m.Group[manager.Id] = manager
			m.clientCount += 1
			m.Lock.Unlock()
			log.Printf("client register success, and manager ID = %s", manager.Id)
		case manager := <-m.UnRegister:
			//注销客户端
			m.Lock.Lock()
			if _, ok := m.Group[manager.Id]; ok {
				//删除分组中客户
				delete(m.Group, manager.Id)
				//客户端数量减1
				m.clientCount -= 1
				log.Printf("client unregister success, and manager ID = %s", manager.Id)
			}
			m.Lock.Unlock()
		case data := <-m.BroadCastMessage:
			//将数据广播给所有客户端
			for _, conn := range m.Group {
				if data.IsBroadCast {
					conn.Message <- data.Message
				} else {
					if inSliceStr(conn.Id, data.ClientIDs) {
						conn.Message <- data.Message
					}
				}

			}
		}
	}
}
