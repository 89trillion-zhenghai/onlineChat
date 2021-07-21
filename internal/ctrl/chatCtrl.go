package ctrl

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"onlineChat/internal/handler"
)
//socket 设置
var (
	upGrader = websocket.Upgrader{
		//
		ReadBufferSize: 1023,
		//
		WriteBufferSize: 1023,
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)


func ChatCtrl(res http.ResponseWriter,req *http.Request)  {
	//将请求升级
	conn, err := upGrader.Upgrade(res,req,nil)
	if err != nil {
		http.NotFound(res, req)
		log.Fatal("无法与",conn.RemoteAddr(),"建立连接")
		return
	}
	values := req.URL.Query()
	name := values.Get("name")
	if len(name) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	handler.OnlineChat(conn,name)
}
