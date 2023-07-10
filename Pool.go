package websocket_manager

import (
	"log"
	"sync"
)

/*
 * todo: 这个库中的最高概念，连接池，一个连接池中可以存在多个manager, 一个manager可以链接多个conn
 * author: joker
 * date: 2023年7月10日12:17:25
 * example: 参考一个游戏里有多个房间，多个房间中存在多个玩家， pool对应整个游戏的连接池，房间对应manager, 玩家对应conn
 */

// pool for all connects

type Pool struct {
	Group                map[string]*Manager
	Lock                 sync.Mutex
	Register, UnRegister chan *Manager
	Close                chan int
	clientCount          uint //分组及客户端数量
}

func (p Pool) Run() {
	for {
		select {
		case num := <-p.Close:
			log.Printf("pool close signal = %d", num)
			break
		case manager := <-p.Register:
			//注册客户端
			p.Lock.Lock()
			p.Group[manager.Id] = manager
			p.clientCount += 1
			p.Lock.Unlock()
			log.Printf("manager register success, and manager ID = %s", manager.Id)
		case manager := <-p.UnRegister:
			//注销客户端
			p.Lock.Lock()
			if _, ok := p.Group[manager.Id]; ok {
				//删除分组中客户
				delete(p.Group, manager.Id)
				//客户端数量减1
				p.clientCount -= 1
				log.Printf("manager unregister success, and manager ID = %s", manager.Id)
			}
			p.Lock.Unlock()
		}
	}
}
