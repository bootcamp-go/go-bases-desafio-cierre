package service

import (
	"fmt"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/storage"
)

// NewServiceTicketDefault returns a new instance of default implementation of ServiceTicket
func NewServiceTicketDefault(st storage.StorageTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{st: st}
}

// Periods
var (
	PeriodEarlyMorning = time.Date(0, 0, 0, 6, 0, 0, 0, time.UTC)
	PeriodMorning      = time.Date(0, 0, 0, 12, 0, 0, 0, time.UTC)
	PeriodAfternoon    = time.Date(0, 0, 0, 19, 0, 0, 0, time.UTC)
	PeriodNight        = time.Date(0, 0, 0, 23, 59, 59, 0, time.UTC)
)

// ServiceTicketDefault is a struct that represents a service for tickets
type ServiceTicketDefault struct {
	st storage.StorageTicket
}

// GetTotalTickets returns the total number of tickets by destination
func (s *ServiceTicketDefault) GetTotalTickets(destination string) (total int, err error) {
	// get all tickets
	tickets, err := s.st.ReadAll()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrServiceTicketsInternal, err)
		return
	}

	// count tickets by destination
	for _, t := range tickets {
		if t.Destination == destination {
			total++
		}
	}

	return
}

// GetCountByPeriod returns the total number of tickets by period
func (s *ServiceTicketDefault) GetCountByPeriod(period time.Time) (total int, err error) {
	// check period
	if period != PeriodEarlyMorning && period != PeriodMorning && period != PeriodAfternoon && period != PeriodNight {
		err = ErrServiceTicketsInvalidPeriod
		return
	}

	// get all tickets
	tickets, err := s.st.ReadAll()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrServiceTicketsInternal, err)
		return
	}

	// count tickets by period
	switch period {
	case PeriodEarlyMorning:
		for _, t := range tickets {
			if t.Date.Before(PeriodEarlyMorning) {
				total++
			}
		}
	case PeriodMorning:
		for _, t := range tickets {
			if t.Date.After(PeriodEarlyMorning) && t.Date.Before(PeriodMorning) {
				total++
			}
		}
	case PeriodAfternoon:
		for _, t := range tickets {
			if t.Date.After(PeriodMorning) && t.Date.Before(PeriodAfternoon) {
				total++
			}
		}
	case PeriodNight:
		for _, t := range tickets {
			if t.Date.After(PeriodAfternoon) {
				total++
			}
		}
	}

	return
}

// PercentageByDestination returns the percentage of tickets by destination
func (s *ServiceTicketDefault) PercentageByDestination(destination string) (percentage float64, err error) {
	// get all tickets
	tickets, err := s.st.ReadAll()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrServiceTicketsInternal, err)
		return
	}

	// check if tickets exist
	if len(tickets) == 0 {
		err = ErrServiceTicketsNoTickets
		return
	}

	// count tickets by destination
	var total int
	for _, t := range tickets {
		if t.Destination == destination {
			total++
		}
	}

	// calculate percentage
	percentage = float64(total) / float64(len(tickets))

	return
}