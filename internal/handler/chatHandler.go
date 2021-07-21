package handler

import (
	"github.com/gorilla/websocket"
	"onlineChat/internal/model"
)

var manager = model.GetManager()


func OnlineChat(conn *websocket.Conn,name string)  {
	client := model.GetInstanceOfClient(conn,name)
	manager.Connection(client)
	go client.Send()
	go client.Accept(manager)
}
