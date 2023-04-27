package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/hsmtkk/literate-barnacle/proto/booking"
	"google.golang.org/grpc"
)

type bookingServer struct {
	booking.UnimplementedBookingServiceServer
}

func (s *bookingServer) DummyList(ctx context.Context, in *booking.DummyListRequest) (*booking.DummyListResponse, error) {
	bookings := []*booking.Booking{
		{Id: 1, UserId: 1, RoomId: 1, StartDateTime: 1, EndDateTime: 1},
		{Id: 2, UserId: 2, RoomId: 2, StartDateTime: 2, EndDateTime: 2},
	}
	return &booking.DummyListResponse{Bookings: bookings}, nil
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
	booking.RegisterBookingServiceServer(s, &bookingServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
