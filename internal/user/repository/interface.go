package repository

import (
	userModel "go-adr-clean/internal/user/model"
)

type UserRepo interface {
	List() (data map[int]userModel.User)
	Create(user userModel.User) (data userModel.User)
}
