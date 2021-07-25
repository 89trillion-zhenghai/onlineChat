package main

import (
	"github.com/gin-gonic/gin"
	"onlineChat/app/http"
	"onlineChat/app/ws"
	"onlineChat/logUtil"
)


func main() {
	logUtil.Init()
	ws.Init()
	r := gin.Default()
	http.Init(r)
	r.Run(":8082")
}