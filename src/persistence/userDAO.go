package persistence

import "gopetstore/src/domain"

func FindUserByUserNameAndPassword(u *domain.User) *domain.User {
	// 查询数据库
	if u.UserName == "admin" && u.Password == "123" {
		return u
	} else {
		return nil
	}
}
