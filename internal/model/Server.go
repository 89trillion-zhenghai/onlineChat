package model

var manager *ClientManager

type ClientManager struct {
	Clients		map[*Client]bool	//统一管理连接的客户端，在线true，下线false
	Message		chan []byte			//接收用户发送的信息
	Register	chan *Client		//新连接的用户
	UnRegister	chan *Client		//断开连接的用户
}

func NewCManager() *ClientManager{
	if manager == nil{
		manager = &ClientManager{
			Clients: make(map[*Client]bool),
			Message: make(chan []byte),
			Register: make(chan *Client),
			UnRegister: make(chan *Client),
		}
	}
	return manager
}

//GetUserList 获取在线用户列表
func (manager *ClientManager) GetUserList() []string{
	userList := make([]string,0)
	for c,flag := range manager.Clients {
		if flag{
			userList = append(userList, c.Name)
		}
	}
	return userList
}