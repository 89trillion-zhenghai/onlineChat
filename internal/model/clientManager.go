package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"onlineChat/message"
	"strings"
)
var manager = GetInstanceOfManager()

func GetManager() *ClientManager {
	if manager == nil{
		return GetInstanceOfManager()
	}
	return manager
}

type ClientManager struct {
	clients		map[*Client]bool	//统一管理连接的客户端，在线true，下线false
	message		chan []byte			//接收用户发送的信息
	register	chan *Client		//新连接的用户
	unRegister	chan *Client		//断开连接的用户
}

type Client struct {
	name	string			//客户端标识符
	socket	*websocket.Conn	//连接对象
	sendMsg	chan []byte		//待发送的信息
}

func GetInstanceOfClient(conn *websocket.Conn,name string) *Client {
	return &Client{
		name: name,
		socket: conn,
		sendMsg: make(chan []byte),
	}
}

func (manager *ClientManager) Connection(client *Client)  {
	manager.register <- client
}


//GetInstanceOfManager 获取一个客户端管理对象
func GetInstanceOfManager() *ClientManager {
	return &ClientManager{
		clients: make(map[*Client]bool),
		message: make(chan []byte),
		register: make(chan *Client),
		unRegister: make(chan *Client),
	}
}

//Start 服务器启动
func (manager *ClientManager) Start()  {
	fmt.Println("server start...")
	for {
		select  {
		case conn := <-manager.register:
			manager.clients[conn] = true
		case conn := <-manager.unRegister:
			//连接断开，清理用户连接信息，关闭接收信息的通道
			if _,flag := manager.clients[conn]; flag{
				close(conn.sendMsg)
				delete(manager.clients,conn)
				conn.socket.Close()
			}
		case msgBytes := <- manager.message:
			msg := &message.Message{}
			proto.Unmarshal(msgBytes,msg)
			switch msg.MsgType {
			case "talk": manager.broadcast(msgBytes)
			case "exit": manager.closeClient(msg.SendName)
			case "userList": manager.sendUserList(msg.SendName)
			case "ping": manager.ping()
			case "pong": manager.pong()
			default:
				fmt.Println("不是规定类型的消息")
			}
		}
	}
}
func (manager *ClientManager) ping()  {

}

func (manager *ClientManager) pong()  {

}

//Broadcast 转发信息给所有人
func (manager *ClientManager) broadcast(msgBytes []byte)  {
	for conn, flag := range manager.clients {
		if flag{
			conn.sendMsg <- msgBytes
		}
	}
}

//断开用户连接
func (manager *ClientManager) closeClient(sendName string)  {
	for conn, flag := range manager.clients {
		if flag{
			if conn.name == sendName{
				manager.unRegister <- conn
			}
		}
	}
}
//发送用户列表信息
func (manager *ClientManager) sendUserList(sendName string)  {
	var client = &Client{}
	msg := &message.Message{}
	for conn, flag := range manager.clients {
		msg.UserList = append(msg.UserList, conn.name)
		if flag{
			if conn.name == sendName{
				client = conn
			}
		}
	}
	bytes, err := proto.Marshal(msg)
	if err != nil{
		log.Fatal(err.Error())
	}
	client.sendMsg <- bytes
}


//Send 给客户端发送信息
func (conn *Client) Send()  {
	defer func() {
		conn.socket.Close()
	}()
	for {
		select {
		case msg := <- conn.sendMsg:
			conn.socket.WriteMessage(websocket.TextMessage,msg)
		}
	}
}

//Accept 接收客户端发送的信息
func (conn *Client) Accept(manager *ClientManager)  {
	defer func() {
		conn.socket.Close()
	}()
	for {
		_, msg, err := conn.socket.ReadMessage()
		//如果出错就让该用户下线
		if err != nil{
			manager.unRegister <- conn
			conn.socket.Close()
			break
		}
		message := stringToProto(string(msg), conn.name)
		manager.message <- message
	}
}
func stringToProto(msg string,name string) []byte{
	msgStr := strings.Split(msg, "_")
	message := &message.Message{
		MsgType: msgStr[0],
		MsgContent: msgStr[1],
		SendName: name,
		UserList: make([]string,0),
	}
	resByte, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("解析失败！格式错误")
	}
	return resByte
}
