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
	err := persistence.InsertAccount(account)
	if err != nil {
		return err
	}
	err = persistence.InsertProfile(account)
	if err != nil {
		return err
	}
	err = persistence.InsertSigOn(account.UserName, account.Password)
	if err != nil {
		return err
	}
	return nil
}

// update account by userName
func UpdateAccount(account *domain.Account) error {
	err := persistence.UpdateAccountByUserName(account, account.UserName)
	if err != nil {
		return err
	}
	err = persistence.UpdateProfileByUserName(account, account.UserName)
	if err != nil {
		return err
	}
	if len(account.Password) > 0 {
		err = persistence.UpdateSignOn(account.UserName, account.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
