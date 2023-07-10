package websocket_manager

/*
 * date: 2023年7月10日12:21:59
 * author: joker
 * todo: 完成管理器的初始化
 */

func InitializationPool() *Pool {
	// todo: new pool
	newPool := new(Pool)

	// init
	newPool.Group = make(map[string]*Manager)
	newPool.Register = make(chan *Manager)
	newPool.UnRegister = make(chan *Manager)
	newPool.Close = make(chan int)

	// 调用监听队列

	go newPool.Run()

	return newPool
}

func InitializationManager(ID string) *Manager {
	// todo: new manager
	newManager := new(Manager)

	// init
	newManager.Id = ID
	newManager.Group = make(map[string]*Client)
	newManager.Register = make(chan *Client)
	newManager.UnRegister = make(chan *Client)
	newManager.BroadCastMessage = make(chan *BroadCastMessageData)
	newManager.Close = make(chan int)

	return newManager
}
