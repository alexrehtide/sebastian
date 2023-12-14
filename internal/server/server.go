package server

import (
	"fmt"
	"net"

	pb "github.com/alexrehtide/sebastian/api/v1"

	"google.golang.org/grpc"
)

func NewServer() *Server {
	s := grpc.NewServer()
	pb.RegisterSebastianServer(s, &serverImpl{})

	return &Server{
		grpcServer: s,
	}
}

type Server struct {
	grpcServer *grpc.Server
}

func (s *Server) Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	return s.grpcServer.Serve(lis)
}
