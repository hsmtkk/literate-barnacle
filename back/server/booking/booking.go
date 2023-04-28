package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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
		{Id: 1, UserId: 1, RoomId: 1, StartDateTime: "2023-04-28T15:19:46", EndDateTime: "2023-04-28T15:19:46"},
		{Id: 2, UserId: 2, RoomId: 2, StartDateTime: "2023-04-28T15:19:46", EndDateTime: "2023-04-28T15:19:46"},
	}
	return &booking.DummyListResponse{Bookings: bookings}, nil
}

func (s *bookingServer) Create(ctx context.Context, in *booking.CreateReqeust) (*booking.CreateResponse, error) {
	id := rand.Int63()
	userID := in.Booking.GetUserId()
	roomID := in.Booking.GetRoomId()
	startDateTime := in.Booking.GetStartDateTime()
	endDateTime := in.Booking.GetEndDateTime()
	newBooking := booking.Booking{Id: id, UserId: userID, RoomId: roomID, StartDateTime: startDateTime, EndDateTime: endDateTime}
	return &booking.CreateResponse{Booking: &newBooking}, nil
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
