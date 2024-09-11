package usecase

import userModel "go-adr-clean/internal/user/model"

func (u *userUsecase) ListAllUser() (data map[int]userModel.User) {
	data = u.UserRepo.List()
	return
}

func (u *userUsecase) CreateUser(user userModel.User) (data userModel.User) {
	data = u.UserRepo.Create(user)
	return
}
