package http

import (
	delivery "go-adr-clean/internal/user/delivery"
	userUsecase "go-adr-clean/internal/user/usecase"
	"net/http"
)

type UserHttpHandler struct {
	Usecase userUsecase.UserUsecase
}

func NewUserHttpHandler(mux *http.ServeMux, usecase userUsecase.UserUsecase) delivery.UserDelivery {
	handler := &UserHttpHandler{
		Usecase: usecase,
	}
	mux.HandleFunc("/api/v1/users", handler.GetAllUsers)
	mux.HandleFunc("/api/v1/user/create", handler.CreateUser)

	return handler

}
