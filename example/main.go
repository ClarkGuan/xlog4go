package main

import (
	logger "github.com/ClarkGuan/xlog4go"

	"time"
)

func main() {
	if err := logger.SetupLogWithFile("./log.json"); err != nil {
		panic(err)
	}
	defer logger.Close()

	var name = "shengkehua"
	for {
		logger.Trace("log4go by %s", name)
		logger.Debug("log4go by %s", name)
		logger.Info("log4go by %s", name)
		logger.Warn("log4go by %s", name)
		logger.Error("log4go by %s", name)
		logger.Fatal("log4go by %s", name)

		time.Sleep(time.Second * 1)
	}
}
