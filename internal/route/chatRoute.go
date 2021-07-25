package route

import (
	"github.com/gin-gonic/gin"
	"onlineChat/internal/ctrl"
)

func Route(r *gin.Engine)  {
	r.GET("/ws",ctrl.ChatCtrl)
}