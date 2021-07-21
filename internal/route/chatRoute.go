package route

import (
	"net/http"
	"onlineChat/internal/ctrl"
)

func Route()  {
	http.HandleFunc("/",ctrl.ChatCtrl)
}