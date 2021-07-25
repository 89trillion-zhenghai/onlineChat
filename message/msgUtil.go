package message

import (
	"google.golang.org/protobuf/proto"
	"onlineChat/logUtil"
)

func SendUserList(list [] string) []byte{
	msg := Message{
		MsgType: "userList",
		UserList: list,
	}
	sendMsg, err := proto.Marshal(&msg)
	if err != nil {
		logUtil.Log.Printf("proto Marshal %s",err.Error())
	}
	return sendMsg
}

func SendTalk(content string) []byte{
	msg := Message{
		MsgType: "talk",
		MsgContent: content,
	}
	sendMsg, err := proto.Marshal(&msg)
	if err != nil {
		logUtil.Log.Printf("proto Marshal %s",err.Error())
	}
	return sendMsg
}
