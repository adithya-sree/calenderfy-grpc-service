package app

import (
	"calenderfy-grpc-service/app/config"
	"calenderfy-grpc-service/app/database"
	"calenderfy-grpc-service/app/logger"
	"calenderfy-grpc-service/app/server"
	pb "calenderfy-grpc-service/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var out *log.Logger = logger.GetLogger("app.go")

type App struct {
	server *server.CalenderfyServer
	config config.Configs
}

func NewApp(c config.Configs) (*App, error) {
	db, err := database.NewDatabase(c)
	if err != nil {
		out.Println("error while initialzing application", err)
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
	pb.RegisterCalenderfyServer(grpcServer, a.server)
	grpcServer.Serve(lis)
}
