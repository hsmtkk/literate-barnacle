package main

import (
	"context"
	"log"
	"time"

	"github.com/hsmtkk/literate-barnacle/proto/booking"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "127.0.0.1:8000"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := booking.NewBookingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DummyList(ctx, &booking.DummyListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", r)
}
