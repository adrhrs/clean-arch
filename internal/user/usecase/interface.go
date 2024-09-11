package usecase

import (
	userModel "go-adr-clean/internal/user/model"
)

type UserUsecase interface {
	ListAllUser() (data map[int]userModel.User)
	CreateUser(user userModel.User) (data userModel.User)
}
