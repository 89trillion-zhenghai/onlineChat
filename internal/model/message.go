package model

type Message struct {
	msgType 	string		//消息类型
	msgContent 	string		//消息内容
	sendName	string		//发送方名字
	userList	[]string	//用户列表
}