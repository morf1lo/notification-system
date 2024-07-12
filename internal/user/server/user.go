package server

import (
	"context"

	"github.com/morf1lo/notification-system/internal/user/model"
	"github.com/morf1lo/notification-system/internal/user/pb"
)

func (s *GRPCServer) Subscribe(ctx context.Context, req *pb.SubscribeReq) (*pb.Empty, error) {
	sub := model.Subscriber{
		Email: req.GetEmail(),
	}
	err := s.services.User.Subscribe(ctx, &sub)
	return &pb.Empty{}, err
}

func (s *GRPCServer) GetAllSubscribers(ctx context.Context, in *pb.Empty) (*pb.GetAllSubscribersRes, error) {
	subs, err := s.services.User.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetAllSubscribersRes{
		Subs: subs,
	}, nil
}
