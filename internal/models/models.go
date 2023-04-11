package models

import "time"


// Use  is the users model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time  
}

// Rooms is the rooms model
type Room struct {

	ID          int
	RoomName    string
	CreatedAt   time.Time
	UpdatedAt   time.Time  
}

// restrictions is the restrictions model
type Restriction struct {
	ID               int
	RestrictionName  string
	CreatedAt        time.Time
	UpdatedAt        time.Time            
}

// Reservatios is the reservations model
type Reservation struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	StartDate   time.Time
	EndDate     time.Time  
	RoomId      int
	CreatedAt   time.Time
	UpdatedAt   time.Time  
	Room        Room
}

// RoomRestrictions is the room restrictions model
type RoomRestrictions struct {
	ID            int
	StartDate      time.Time
	EndDate        time.Time  
	RoomId         int
	ReservationId  int
	RestrictionsId int 
	CreatedAt      time.Time
	UpdatedAt      time.Time  
	Room           Room
	Reservation    Reservation
	RestrictionId  Restriction
}