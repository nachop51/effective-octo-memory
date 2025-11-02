package accounts

type AccountInterface interface {
	GetAccounts() ([]*Account, error)
	CreateAccount(account *Account) error
}
