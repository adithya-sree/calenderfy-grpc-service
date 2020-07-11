package main

import (
	"calenderfy-grpc-service/app"
	"calenderfy-grpc-service/app/config"
	"calenderfy-grpc-service/app/logger"
	"log"
)

var out *log.Logger = logger.GetLogger("main.go")
var c config.Configs = config.GetConfigs()

func main() {
	out.Println("starting calenderfy-grpc-service")

	app, err := app.NewApp(c)
	if err != nil {
		out.Println("service was not started, exceptions during startup", err)
		return
	}

	app.Run()
}
