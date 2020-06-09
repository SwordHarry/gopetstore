package service

import (
	"gopetstore/src/domain"
	"gopetstore/src/persistence"
)

func Login(u *domain.User) *domain.User {
	// 登录业务相关代码
	return persistence.FindUserByUserNameAndPassword(u)
}
