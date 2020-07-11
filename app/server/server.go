package server

import (
	"calenderfy-grpc-service/app/dao"
	"calenderfy-grpc-service/app/logger"
	pb "calenderfy-grpc-service/proto"
	"context"
	"log"
	"time"
)

var startTime time.Time
var out *log.Logger = logger.GetLogger("server.go")

func init() {
	startTime = time.Now()
}

type CalenderfyServer struct {
	pb.UnimplementedCalenderfyServer
	db dao.Dao
}

func NewServer(dao dao.Dao) *CalenderfyServer {
	return &CalenderfyServer{db: dao}
}

func (s CalenderfyServer) Running(ctx context.Context, in *pb.EmptyRequest) (*pb.RunningResponse, error) {
	out.Println("running check received")
	return &pb.RunningResponse{Running: true}, nil
}

func (s CalenderfyServer) Uptime(ctx context.Context, in *pb.EmptyRequest) (*pb.UptimeResponse, error) {
	out.Println("uptime check received")
	return &pb.UptimeResponse{StartedAt: startTime.Format("2006.01.02 15:04:05"), Uptime: time.Since(startTime).Milliseconds()}, nil
}
