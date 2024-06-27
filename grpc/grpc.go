package grpc

import (
	"go_schedule_service/config"
	"go_schedule_service/genproto/lesson_service"

	"go_schedule_service/grpc/client"
	"go_schedule_service/grpc/service"
	"go_schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	lesson_service.RegisterLessonServiceServer(grpcServer, service.NewLessonService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
