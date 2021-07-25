package handler

import (
	"github.com/gorilla/websocket"
	"onlineChat/app/ws"
	"onlineChat/internal/model"
)

func OnlineChat(conn *websocket.Conn,name string)  {
	client := &model.Client{
		Manager: model.NewCManager(),
		Name: name,
		Socket: conn,
		SendMsg: make(chan []byte,256),
	}
	client.CreateClient()
	go ws.AcceptMsg(client)
	go ws.SendMsg(client)
}
