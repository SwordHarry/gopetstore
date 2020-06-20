package persistence

import (
	"gopetstore/src/domain"
	"testing"
)

func TestInsertAccount(t *testing.T) {
	err := InsertAccount(&domain.Account{
		UserName:            "1234",
		Password:            "1234",
		Email:               "1234",
		FirstName:           "1234",
		LastName:            "1234",
		Status:              "",
		Address1:            "1234",
		Address2:            "1234",
		City:                "1234",
		State:               "1234",
		Zip:                 "1234",
		Country:             "1234",
		Phone:               "1234",
		FavouriteCategoryId: "1234",
		LanguagePreference:  "1234",
		ListOption:          false,
		BannerOption:        false,
		BannerName:          "",
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateAccountByUserName(t *testing.T) {
	err := UpdateAccountByUserName(&domain.Account{
		UserName:            "test",
		Password:            "hahaha",
		Email:               "hahaha",
		FirstName:           "test",
		LastName:            "test",
		Status:              "",
		Address1:            "test",
		Address2:            "test",
		City:                "test",
		State:               "test",
		Zip:                 "test",
		Country:             "test",
		Phone:               "1234",
		FavouriteCategoryId: "test",
		LanguagePreference:  "test",
		ListOption:          true,
		BannerOption:        true,
		BannerName:          "",
	}, "test")
	if err != nil {
		t.Error(err.Error())
	}
}
