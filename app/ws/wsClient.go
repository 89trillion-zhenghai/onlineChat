package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"onlineChat/internal/model"
	"onlineChat/logUtil"
	Message "onlineChat/message"
	"time"
)
const (
	writeWait = 10 * time.Second
	pongWait = 10 * time.Second
	pingPeriod = (pongWait * 9) / 10
)


//SendMsg 发送消息
func SendMsg(c *model.Client)  {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Socket.Close()
	}()
	for {
		select {
		case message := <-c.SendMsg:
			c.Socket.WriteMessage(websocket.TextMessage,message)
		case <- ticker.C:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.Socket.WriteMessage(websocket.PingMessage, nil)
			if err != nil{
				return
			}
		}
	}
}

//AcceptMsg 接收消息
func AcceptMsg(c *model.Client) {
	defer func() {
		c.Manager.UnRegister <- c
		c.Socket.Close()
	}()
	c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	c.Socket.SetPongHandler(func(string) error {
		c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		fmt.Println("心跳检测"+time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	for{
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			break
		}
		message := Message.Message{}
		err = proto.Unmarshal(msg, &message)
		if err != nil{
			logUtil.Log.Printf("error: %v", err)
		}
		switch message.MsgType {
		case "exit":
			break
		case "userList":
			userList := c.Manager.GetUserList()
			c.SendMsg <- Message.SendUserList(userList)
		default:
			c.Manager.Message <- msg
		}
	}
}

