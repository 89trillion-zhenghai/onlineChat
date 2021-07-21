package ws

import "onlineChat/internal/model"

func Init()  {
	manager := model.GetManager()
	go manager.Start()
}

