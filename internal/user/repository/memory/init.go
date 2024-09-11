package memory

import (
	userModel "go-adr-clean/internal/user/model"
	repo "go-adr-clean/internal/user/repository"
)

type userMemoryRepo struct {
	userData map[int]userModel.User
}

func NewUserMemoryRepo(userData map[int]userModel.User) repo.UserRepo {
	return &userMemoryRepo{
		userData: userData,
	}
}
