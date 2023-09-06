package service

import (
	"errors"
	"time"
)

// ServiceTicket is an interface that represents a service for tickets
type ServiceTicket interface {
	// GetTotalTickets returns the total number of tickets by destination
	GetTotalTickets(destination string) (total int, err error)

	// GetCountByPeriod returns the total number of tickets by period
	GetCountByPeriod(period time.Time) (total int, err error)

	// PercentageByDestination returns the percentage of tickets by destination
	PercentageByDestination(destination string) (percentage float64, err error)
}

var (
	ErrServiceTicketsInternal	   = errors.New("internal service error")
	ErrServiceTicketsInvalidPeriod = errors.New("invalid period")
	ErrServiceTicketsNoTickets	   = errors.New("no tickets")
)