package models

import (
	"time"
)

// User is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

type RoomRestriction struct {
	ID            int
	RoomID        int
	ReservationID int
	RestrictionID int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}

// MailData holds an email message
type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
}
