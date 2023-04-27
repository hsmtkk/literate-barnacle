package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/hsmtkk/literate-barnacle/proto/user"
	"google.golang.org/grpc"
)

type userServer struct {
	user.UnimplementedUserServiceServer
}

func (s *userServer) DummyList(ctx context.Context, in *user.DummyListRequest) (*user.DummyListResponse, error) {
	users := []*user.User{
		{Id: 1, Name: "dummy1"},
		{Id: 2, Name: "dummy2"},
	}
	return &user.DummyListResponse{Users: users}, nil
}

func main() {
	portStr := os.Getenv(("PORT"))
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("failed to parse %s as int: %v", portStr, err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &userServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
