package storage

// NewStorageTicketMock returns a new instance of mock implementation of StorageTicket
func NewStorageTicketMock() *StorageTicketMock {
	return &StorageTicketMock{}
}

// Ticket is a struct that represents a ticket
type StorageTicketMock struct {
	// FuncReadAll is a function that mocks the ReadAll method
	FuncReadAll func() (t []*Ticket, err error)
}

// ReadAll returns all tickets
func (s *StorageTicketMock) ReadAll() (t []*Ticket, err error) {
	t, err = s.FuncReadAll()
	return
}
