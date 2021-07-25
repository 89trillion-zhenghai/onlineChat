package ws

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"onlineChat/internal/model"
	"onlineChat/logUtil"
	"onlineChat/message"
)

func Init(){
	manager := model.NewCManager()
	go Start(manager)
}

func Start(manager *model.ClientManager)  {
	logUtil.Log.Printf("Server start....")
	fmt.Println("Server start...")
	for{
		select {
		case client := <-manager.Register:
			manager.Clients[client] = true
			logUtil.Log.Printf("%s上线\n",client.Name)
			msgByte := message.SendTalk(client.Name + "\tis connected")
			Broadcast(msgByte,manager)
		case client := <-manager.UnRegister:
			if _, ok := manager.Clients[client]; ok {
				delete(manager.Clients, client)
				close(client.SendMsg)
				logUtil.Log.Printf("%s下线\n",client.Name)
			}
			msgByte := message.SendTalk(client.Name + "\tis Disconnected")
			Broadcast(msgByte,manager)
		case acceptMsg := <-manager.Message:
			msg := message.Message{}
			proto.Unmarshal(acceptMsg,&msg)
			switch msg.MsgType {
			case "talk": Broadcast(acceptMsg,manager)
			default:
				fmt.Println("不是规定类型的消息")
			}
		}
	}
}

//Broadcast 广播信息
func Broadcast(msgBytes []byte, manager *model.ClientManager)  {
	for conn, flag := range manager.Clients {
		if flag{
			conn.SendMsg <- msgBytes
		}
	}
}

