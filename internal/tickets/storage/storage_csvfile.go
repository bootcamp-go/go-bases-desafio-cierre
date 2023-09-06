package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// NewStorageTicketCSVFile returns a new instance of StorageTicketCSVFile
func NewStorageTicketCSVFile(path string) *StorageTicketCSVFile {
	return &StorageTicketCSVFile{Path: path}
}

// StorageTicketCSVFile is a struct that represents a storage for tickets in a CSV file
type StorageTicketCSVFile struct {
	Path string
}

// GetTotalTickets returns the total number of tickets by destination
func (s *StorageTicketCSVFile) ReadAll() (t []*Ticket, err error) {
	// open file
	f, err := os.Open(s.Path)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageTicketsInternal, err.Error())
		return
	}
	defer f.Close()

	// reader
	reader := csv.NewReader(f)
	for {
		// read line
		var records []string
		records, err = reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
				break
			} else {
				err = fmt.Errorf("%w. %s", ErrStorageTicketsInternal, err.Error())
				return
			}
		}

		// serialization
		var ticket Ticket
		// -> id
		ticket.Id, err = strconv.Atoi(records[0])
		if err != nil {
			err = fmt.Errorf("%w. %s", ErrStorageTicketsInternal, err.Error())
			return
		}
		// -> name
		ticket.Name = records[1]
		// -> email
		ticket.Email = records[2]
		// -> destination
		ticket.Destination = records[3]

		// -> date
		var hour, minute int
		period := strings.Split(records[4], ":")
		hour, err = strconv.Atoi(period[0])
		if err != nil {
			err = fmt.Errorf("%w. %s", ErrStorageTicketsInternal, err.Error())
			return
		}
		minute, err = strconv.Atoi(period[1])
		if err != nil {
			err = fmt.Errorf("%w. %s", ErrStorageTicketsInternal, err.Error())
			return
		}

		ticket.Date = time.Date(0, 0, 0, hour, minute, 0, 0, time.UTC)

		// -> price
		ticket.Price, err = strconv.ParseFloat(records[5], 64)
		if err != nil {
			err = fmt.Errorf("%w. %s", ErrStorageTicketsInternal, err.Error())
			return
		}

		// append ticket
		t = append(t, &ticket)
	}

	return
}