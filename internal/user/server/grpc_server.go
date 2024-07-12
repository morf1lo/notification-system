package server

import (
	"net"

	"github.com/morf1lo/notification-system/internal/user/config"
	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/internal/user/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedUserServer

	services *service.Service
}

func Run(cfg *config.GRPCServerConfig, services *service.Service) error {
	lis, err := net.Listen(cfg.Network, cfg.Addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	s := grpc.NewServer(
		grpc.MaxSendMsgSize(1024 * 1024 * 32),
	)
	pb.RegisterUserServer(s, &GRPCServer{services: services})
	return s.Serve(lis)
}
