package http

import (
	"github.com/gin-gonic/gin"
	"onlineChat/internal/route"
)

func Init(r *gin.Engine)  {
	route.Route(r)
	return
}
