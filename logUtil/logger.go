package logUtil

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	logFile, err := os.OpenFile("./log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("can't open server.log,please try again")
	}
	parting := "-----------------------------------------------------------------\n"
	logFile.Write([]byte(parting))
	Log = log.New(logFile,"log:\t",log.Ldate | log.Ltime | log.Lshortfile)
}


