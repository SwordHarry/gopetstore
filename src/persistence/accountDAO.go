package persistence

import (
	"database/sql"
	"errors"
	"gopetstore/src/domain"
	"gopetstore/src/util"
)

// all SQL about Account
// get account by userName from signOn, account, bannerData
const getAccountByUsernameSQL = `SELECT SIGNON.USERNAME,ACCOUNT.EMAIL,ACCOUNT.FIRSTNAME,ACCOUNT.LASTNAME,ACCOUNT.STATUS,ACCOUNT.ADDR1 AS address1,
ACCOUNT.ADDR2 AS address2,ACCOUNT.CITY,ACCOUNT.STATE,ACCOUNT.ZIP,ACCOUNT.COUNTRY,ACCOUNT.PHONE,PROFILE.LANGPREF AS languagePreference,
PROFILE.FAVCATEGORY AS favouriteCategoryId,PROFILE.MYLISTOPT AS listOption,PROFILE.BANNEROPT AS bannerOption,BANNERDATA.BANNERNAME 
FROM ACCOUNT, PROFILE, SIGNON, BANNERDATA
WHERE ACCOUNT.USERID = ? AND SIGNON.USERNAME = ACCOUNT.USERID AND PROFILE.USERID = ACCOUNT.USERID AND PROFILE.FAVCATEGORY = BANNERDATA.FAVCATEGORY`

// get account by userName and password from signOn, account, bannerData
const getAccountByUsernameAndPasswordSQL = `SELECT SIGNON.USERNAME,ACCOUNT.EMAIL,ACCOUNT.FIRSTNAME,ACCOUNT.LASTNAME,ACCOUNT.STATUS,ACCOUNT.ADDR1 AS address1,
ACCOUNT.ADDR2 AS address2,ACCOUNT.CITY,ACCOUNT.STATE,ACCOUNT.ZIP,ACCOUNT.COUNTRY,ACCOUNT.PHONE,PROFILE.LANGPREF AS languagePreference,
PROFILE.FAVCATEGORY AS favouriteCategoryId,PROFILE.MYLISTOPT AS listOption,PROFILE.BANNEROPT AS bannerOption,BANNERDATA.BANNERNAME 
FROM ACCOUNT, PROFILE, SIGNON, BANNERDATA WHERE ACCOUNT.USERID = ? AND SIGNON.PASSWORD = ? 
AND SIGNON.USERNAME = ACCOUNT.USERID AND PROFILE.USERID = ACCOUNT.USERID AND PROFILE.FAVCATEGORY = BANNERDATA.FAVCATEGORY`

// update account from account
const updateAccountSQL = `UPDATE ACCOUNT SET EMAIL = ?,FIRSTNAME = ?,LASTNAME = ?,STATUS = ?,ADDR1 = ?,
ADDR2 = ?,CITY = ?,STATE = ?,ZIP = ?,COUNTRY = ?,PHONE = ? WHERE USERID = ?`

// insert account from account
const insertAccountSQL = `INSERT INTO ACCOUNT (EMAIL, FIRSTNAME, LASTNAME, STATUS, ADDR1, ADDR2, CITY, STATE, ZIP, COUNTRY, PHONE, USERID) 
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

// update profile from profile
const updateProfileSQL = `UPDATE PROFILE SET LANGPREF = ?, FAVCATEGORY = ?,mylistopt = ?,banneropt = ? WHERE USERID = ?`

// insert profile from profile
const insertProfileSQL = `INSERT INTO PROFILE (LANGPREF, FAVCATEGORY, USERID, mylistopt, banneropt) VALUES (?, ?, ?, ?, ?)`

// update password by userName from signOn
const updateSigOnSQL = `UPDATE SIGNON SET PASSWORD = ? WHERE USERNAME = ?`

// insert username and password from signOn
const insertSigOnSQL = `INSERT INTO SIGNON (USERNAME,PASSWORD) VALUES (?, ?)`

func scanAccountWithSignOnAndBannerData(r *sql.Rows) (*domain.Account, error) {
	var userName, email, firstName, lastName, status, addr1, addr2, city, state, zip, country, phone string
	var languagePreference, favouriteCategoryId, bannerName string
	var listOption, bannerOption bool

	err := r.Scan(&userName, &email, &firstName, &lastName, &status, &addr1, &addr2, &city, &state, &zip, &country, &phone,
		&languagePreference, &favouriteCategoryId, &listOption, &bannerOption, &bannerName)
	a := &domain.Account{
		UserName:            userName,
		Password:            "",
		Email:               email,
		FirstName:           firstName,
		LastName:            lastName,
		Status:              status,
		Address1:            addr1,
		Address2:            addr2,
		City:                city,
		State:               state,
		Zip:                 zip,
		Country:             country,
		Phone:               phone,
		FavouriteCategoryId: favouriteCategoryId,
		LanguagePreference:  languagePreference,
		ListOption:          listOption,
		BannerOption:        bannerOption,
		BannerName:          bannerName,
	}
	return a, err
}

// get account by userName for register
func GetAccountByUserName(userName string) (*domain.Account, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getAccountByUsernameSQL, userName)
	if err != nil {
		return nil, err
	}
	if r.Next() {
		a, err := scanAccountWithSignOnAndBannerData(r)
		if err != nil {
			return nil, err
		}
		return a, nil
	}
	return nil, errors.New("can not find the account by this user name")
}

// get account by user name and password for signIn
func GetAccountByUserNameAndPassword(userName string, password string) (*domain.Account, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := d.Query(getAccountByUsernameAndPasswordSQL, userName, password)
	if err != nil {
		return nil, err
	}
	if r.Next() {
		a, err := scanAccountWithSignOnAndBannerData(r)
		if err != nil {
			return nil, err
		}
		return a, nil
	}
	return nil, errors.New("can not find the account by this user name and password")
}

// update account by userName
func UpdateAccountByUserName(account *domain.Account, userName string) error {
	return util.InsertOrUpdate(updateAccountSQL, "none account was updated by this userName",
		account.Email, account.FirstName, account.LastName, account.Status, account.Address1, account.Address2,
		account.City, account.State, account.Zip, account.Country, account.Phone, userName)
}

// insert account
func InsertAccount(account *domain.Account) error {
	return util.InsertOrUpdate(insertAccountSQL,
		"insert account error", account.Email, account.FirstName, account.LastName,
		account.Status, account.Address1, account.Address2, account.City, account.State,
		account.Zip, account.Country, account.Phone, account.UserName)
}

// update profile by userName
func UpdateProfileByUserName(account *domain.Account, userName string) error {
	return util.InsertOrUpdate(updateProfileSQL,
		"can not update profile by this user name", account.LanguagePreference, account.FavouriteCategoryId,
		account.ListOption, account.BannerOption, userName)
}

// insert profile from profile
func InsertProfile(account *domain.Account) error {
	return util.InsertOrUpdate(insertProfileSQL, "can not insert profile", account.LanguagePreference,
		account.FavouriteCategoryId, account.UserName, account.ListOption, account.BannerOption)
}

// update userName and password from signOn
func UpdateSignOn(userName string, password string) error {
	return util.InsertOrUpdate(updateSigOnSQL, "can not update password by this user name", password, userName)
}

// insert user name and password from signOn
func InsertSigOn(userName string, password string) error {
	return util.InsertOrUpdate(insertSigOnSQL, "can not insert this user name and password", userName, password)
}
