package server

import (
	"context"

	pb "github.com/alexrehtide/sebastian/api/v1"
)

type serverImpl struct {
	pb.UnimplementedSebastianServer
}

func (s *serverImpl) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	return &pb.AuthenticateResponse{AccessToken: "xxx", RefreshToken: "yyy"}, nil
}
