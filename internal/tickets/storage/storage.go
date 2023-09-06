package storage

import (
	"errors"
	"time"
)

// Ticket is a struct that represents a ticket
type Ticket struct {
	Id			int
	Name		string
	Email		string
	Destination	string
	Date		time.Time
	Price		float64
}

// StorageTicket is an interface that represents a storage for tickets
type StorageTicket interface {
	// ReadAll returns all tickets
	ReadAll() (t []*Ticket, err error)
}

var (
	ErrStorageTicketsInternal = errors.New("internal storage error")
)