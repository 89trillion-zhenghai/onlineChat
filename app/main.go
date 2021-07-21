package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"onlineChat/app/ws"
	"onlineChat/internal/route"
	"onlineChat/message"
	"strings"
)


func main() {
	ws.Init()
	route.Route()
	http.ListenAndServe(":8082",nil)
	//testProto("talk_你好世界")
}

func testProto(str string) {
	//str: talk_content
	msgStr := strings.Split(str, "_")
	msg := &message.Message{
		MsgType: msgStr[0],
		MsgContent: msgStr[1],
		SendName: "smallbai",
		UserList: make([]string,0),
	}
	bytes, err := proto.Marshal(msg)
	if err != nil{
		log.Fatal("解析失败")
	}
	fmt.Println(bytes)
	msg01 := &message.Message{}
	err = proto.Unmarshal(bytes, msg01)
	if err != nil {
		log.Fatal("解析失败")
	}
	fmt.Println(msg01)
}