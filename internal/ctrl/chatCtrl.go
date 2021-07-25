package ctrl

import (
	"github.com/gin-gonic/gin"
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


func ChatCtrl(c *gin.Context)  {
	//将请求升级
	conn, err := upGrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		log.Fatal("无法与",conn.RemoteAddr(),"建立连接")
		return
	}

	name := c.Request.Header.Get("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"无法连接",
		})
		return
	}
	handler.OnlineChat(conn,name)
}
