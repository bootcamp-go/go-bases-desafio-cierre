package service

import (
	"testing"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/storage"
	"github.com/stretchr/testify/assert"
)

// Tests for ServiceTicketDefault
func TestServiceTicketDefault_GetTotalTickets(t *testing.T) {
	t.Run("should return the total number of tickets by destination", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Brazil"
		total, err := sv.GetTotalTickets(inputDestination)

		// assert
		expectedTotal := 1
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return 0, no tickets", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Brazil"
		total, err := sv.GetTotalTickets(inputDestination)

		// assert
		expectedTotal := 0
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return an error when the storage returns an error", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			err = storage.ErrStorageTicketsInternal
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Brazil"
		total, err := sv.GetTotalTickets(inputDestination)

		// assert
		expectedTotal := 0
		expectedErr := ErrServiceTicketsInternal
		expectedErrMsg := "internal service error. internal storage error"
		assert.Equal(t, expectedTotal, total)
		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, expectedErrMsg, err.Error())
	})
}

func TestServiceTicketDefault_GetCountByPeriod(t *testing.T) {
	t.Run("should return 1 ticket for early morning period", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date: 		time.Date(0, 0, 0, 5, 59, 59, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 6, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}

			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := PeriodEarlyMorning
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 1
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return 1 ticket for morning period", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date: 		time.Date(0, 0, 0, 11, 59, 59, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 12, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}

			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := PeriodMorning
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 1
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return 1 ticket for afternoon period", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date: 		time.Date(0, 0, 0, 18, 59, 59, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 19, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}

			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := PeriodAfternoon
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 1
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return 1 ticket for night period", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date: 		time.Date(0, 0, 0, 23, 59, 59, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}

			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := PeriodNight
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 1
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return 0 ", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := PeriodNight
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 0
		expectedErr := error(nil)
		assert.Equal(t, expectedTotal, total)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return an error when the storage returns an error", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			err = storage.ErrStorageTicketsInternal
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := PeriodNight
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 0
		expectedErr := ErrServiceTicketsInternal
		expectedErrMsg := "internal service error. internal storage error"
		assert.Equal(t, expectedTotal, total)
		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, expectedErrMsg, err.Error())
	})

	t.Run("should return an error when the period is invalid", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputPeriod := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
		total, err := sv.GetCountByPeriod(inputPeriod)

		// assert
		expectedTotal := 0
		expectedErr := ErrServiceTicketsInvalidPeriod
		expectedErrMsg := "invalid period"
		assert.Equal(t, expectedTotal, total)
		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, expectedErrMsg, err.Error())
	})
}

func TestServiceTicketDefault_PercentageByDestination(t *testing.T) {
	t.Run("should return the percentage of tickets by destination", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}

			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Brazil"
		percentage, err := sv.PercentageByDestination(inputDestination)

		// assert
		expectedPercentage := 0.5
		expectedErr := error(nil)
		assert.Equal(t, expectedPercentage, percentage)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return 0, no tickets of destination", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			t = []*storage.Ticket{
				{
					Id:			1,
					Name:		"John Doe",
					Email:		"johndoe@gmail.com",
					Destination:"Brazil",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
				{
					Id:			2,
					Name:		"Jane Doe",
					Email:		"janedoe@gmail.com",
					Destination:"Argentina",
					Date:		time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					Price:		100.0,
				},
			}

			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Chile"
		percentage, err := sv.PercentageByDestination(inputDestination)

		// assert
		expectedPercentage := 0.0
		expectedErr := error(nil)
		assert.Equal(t, expectedPercentage, percentage)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return an error when the storage returns an error", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			err = storage.ErrStorageTicketsInternal
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Brazil"
		percentage, err := sv.PercentageByDestination(inputDestination)

		// assert
		expectedPercentage := 0.0
		expectedErr := ErrServiceTicketsInternal
		expectedErrMsg := "internal service error. internal storage error"
		assert.Equal(t, expectedPercentage, percentage)
		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, expectedErrMsg, err.Error())
	})

	t.Run("should return an error when no tickets", func(t *testing.T) {
		// arrange
		st := storage.NewStorageTicketMock()
		st.FuncReadAll = func() (t []*storage.Ticket, err error) {
			return
		}

		sv := NewServiceTicketDefault(st)

		// act
		inputDestination := "Brazil"
		percentage, err := sv.PercentageByDestination(inputDestination)

		// assert
		expectedPercentage := 0.0
		expectedErr := ErrServiceTicketsNoTickets
		expectedErrMsg := "no tickets"
		assert.Equal(t, expectedPercentage, percentage)
		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, expectedErrMsg, err.Error())
	})
}
		