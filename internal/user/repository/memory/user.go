package memory

import (
	userModel "go-adr-clean/internal/user/model"
)

func (r *userMemoryRepo) List() (data map[int]userModel.User) {
	data = r.userData
	return
}

func (r *userMemoryRepo) Create(user userModel.User) (data userModel.User) {
	r.userData[user.ID] = user
	data = user
	return
}
