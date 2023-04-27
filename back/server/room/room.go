package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/hsmtkk/literate-barnacle/proto/room"
	"google.golang.org/grpc"
)

type roomServer struct {
	room.UnimplementedRoomServiceServer
}

func (s *roomServer) DummyList(ctx context.Context, in *room.DummyListRequest) (*room.DummyListResponse, error) {
	rooms := []*room.Room{
		{Id: 1, Name: "dummy1"},
		{Id: 2, Name: "dummy2"},
	}
	return &room.DummyListResponse{Rooms: rooms}, nil
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
	room.RegisterRoomServiceServer(s, &roomServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
