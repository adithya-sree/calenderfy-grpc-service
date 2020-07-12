package server

import (
	"calendarfy-grpc-service/app/dao"
	"calendarfy-grpc-service/app/logger"
	pb "calendarfy-grpc-service/proto"
	"context"
	"time"
)

var startTime time.Time
var out = logger.GetLogger("server.go")

func init() {
	startTime = time.Now()
}

type CalendarfyServer struct {
	pb.UnimplementedCalendarfyServer
	db dao.Dao
}

func NewServer(dao dao.Dao) *CalendarfyServer {
	return &CalendarfyServer{db: dao}
}

func (s CalendarfyServer) Running(ctx context.Context, in *pb.EmptyRequest) (*pb.RunningResponse, error) {
	out.Println("running check received")
	return &pb.RunningResponse{Running: true}, nil
}

func (s CalendarfyServer) Uptime(ctx context.Context, in *pb.EmptyRequest) (*pb.UptimeResponse, error) {
	out.Println("uptime check received")
	return &pb.UptimeResponse{StartedAt: startTime.Format("2006.01.02 15:04:05"), Uptime: time.Since(startTime).Milliseconds()}, nil
}
