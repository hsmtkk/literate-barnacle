package model

import "time"

type Booking struct {
	ID            int
	UserID        int
	RoomID        int
	ReservedNum   int
	StartDateTime time.Time
	EndDateTime   time.Time
}
