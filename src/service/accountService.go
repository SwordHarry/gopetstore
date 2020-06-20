package service

import (
	"gopetstore/src/domain"
	"gopetstore/src/persistence"
)

// get account info by userName for register
func GetAccountByUserName(userName string) (*domain.Account, error) {
	return persistence.GetAccountByUserName(userName)
}

// get account by userName and password for signIn
func GetAccountByUserNameAndPassword(userName string, password string) (*domain.Account, error) {
	return persistence.GetAccountByUserNameAndPassword(userName, password)
}

// insert account
func InsertAccount(account *domain.Account) error {
	return persistence.InsertAccount(account)
}

// update account by userName
func UpdateAccount(account *domain.Account) error {
	return persistence.UpdateAccountByUserName(account, account.UserName)
}
