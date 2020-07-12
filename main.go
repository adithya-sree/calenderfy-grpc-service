package main

import (
	"calendarfy-grpc-service/app"
	"calendarfy-grpc-service/app/config"
	"calendarfy-grpc-service/app/logger"
)

var out = logger.GetLogger("main.go")
var c = config.GetConfigs()

func main() {
	out.Println("starting calendarfy-grpc-service")

	a, err := app.NewApp(c)
	if err != nil {
		out.Println("service was not started, exceptions during startup", err)
		return
	}

	a.Run()
}
