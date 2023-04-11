package models

import "time"

// reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}


// User is the users model
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
type Rooms struct {

	ID          int
	RoomName    string
	CreatedAt   time.Time
	UpdatedAt   time.Time  
}

// restrictions is the restrictions model
type Restrictions struct {
	ID               int
	RestrictionName  string
	CreatedAt        time.Time
	UpdatedAt        time.Time            
}

// Reservatios is the reservations model
type Reservations struct {
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
	Room        Rooms
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
	Room           Rooms
	Reservation    Reservations
	RestrictionId  Restrictions
}