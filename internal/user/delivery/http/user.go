package http

import (
	"encoding/json"
	userModel "go-adr-clean/internal/user/model"
	"log"
	"net/http"
)

func (u *UserHttpHandler) Create() error {
	return http.ErrAbortHandler
}

func (u *UserHttpHandler) List() error {
	return http.ErrAbortHandler
}

func (u *UserHttpHandler) GetAllUsers(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusCreated
		data           map[int]userModel.User
	)

	defer func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	data = u.Usecase.ListAllUser()
}

func (u *UserHttpHandler) CreateUser(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusCreated
		user           userModel.User
	)

	defer func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(user)
	}()

	errDecode := json.NewDecoder(req.Body).Decode(&user)
	if errDecode != nil {
		log.Println(errDecode)
		responseStatus = http.StatusBadRequest
		return
	}
	user = u.Usecase.CreateUser(user)
}
