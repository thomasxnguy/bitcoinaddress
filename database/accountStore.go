package database

import (
	"errors"
	"github.com/google/uuid"
	"github.com/thomasxnguy/bitcoinaddress/models"
)

var (
	ErrUniqueConstraint = errors.New("id already used")
)

// AccountStorer defines database operations for account.
type AccountStorer interface {
	Create(*models.Account) error
	Get(id uuid.UUID) (*models.Account, error)
	Delete(id uuid.UUID)
}

// MockAccountStore implements a mock of AccountStore using in memory as DB,
type MockAccountStore struct {
	Db map[uuid.UUID]*models.Account
}

// NewMockAccountStore returns an AccountStore.
func NewMockAccountStore() AccountStorer {
	db := make(map[uuid.UUID]*models.Account)
	return &MockAccountStore{
		Db: db,
	}
}

// Create creates a new account.
func (s *MockAccountStore) Create(a *models.Account) error {
	if _, ok := s.Db[a.Id]; ok {
		return ErrUniqueConstraint
	}

	s.Db[a.Id] = a
	return nil
}

// Get a account by Id.
func (s *MockAccountStore) Get(id uuid.UUID) (*models.Account, error) {
	return s.Db[id], nil
}

// Delete an account.
func (s *MockAccountStore) Delete(id uuid.UUID) {
	delete(s.Db, id)
}
