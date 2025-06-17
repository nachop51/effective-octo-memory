package accounts

type AccountService struct {
	store AccountInterface
}

func NewAccountService(store AccountInterface) *AccountService {
	return &AccountService{
		store: store,
	}
}

func (s *AccountService) GetAccounts() ([]*Account, error) {
	return s.store.GetAccounts()
}

func (s *AccountService) CreateAccount(body *CreateAccountBody) (*Account, error) {
	account := &Account{
		Name:   body.Name,
		UserID: body.UserID,
	}

	err := s.store.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
