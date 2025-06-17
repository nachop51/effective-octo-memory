package accounts

import "gorm.io/gorm"

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

func (s *AccountStore) GetAccounts() ([]*Account, error) {
	var accounts []*Account

	res := s.db.Find(&accounts)

	if res.Error != nil {
		return nil, res.Error
	}

	return accounts, nil
}

func (s *AccountStore) CreateAccount(account *Account) error {
	return s.db.Create(account).Error
}
