package app

import (
	"calendarfy-grpc-service/app/config"
	"calendarfy-grpc-service/app/dao"
	"calendarfy-grpc-service/app/logger"
	"calendarfy-grpc-service/app/server"
	pb "calendarfy-grpc-service/proto"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var out = logger.GetLogger("app.go")

type App struct {
	server *server.CalendarfyServer
	config config.Configs
}

func NewApp(c config.Configs) (*App, error) {
	db, err := dao.NewDao(c)
	if err != nil {
		out.Println("error while initializing application", err)
		return &App{}, err
	}

	return &App{
		config: c,
		server: server.NewServer(db),
	}, nil
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", a.config.Port))
	if err != nil {
		out.Printf("failed to listen: %v", err)
		return
	}

	out.Println("application starting on port", a.config.Port)

	grpcServer := grpc.NewServer()
	pb.RegisterCalendarfyServer(grpcServer, a.server)
	grpcServer.Serve(lis)
}
