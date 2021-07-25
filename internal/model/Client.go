package model

import "github.com/gorilla/websocket"

type Client struct {
	Manager *ClientManager	//客户端管理
	Name	string			//客户端标识符
	Socket	*websocket.Conn	//连接对象
	SendMsg	chan []byte		//待发送的信息
}

//CreateClient 用户上线
func (client *Client) CreateClient() {
	client.Manager.Clients[client] = true
	client.Manager.Register <- client
}
