package main

import (
	userDeliveryHttp "go-adr-clean/internal/user/delivery/http"
	userModel "go-adr-clean/internal/user/model"
	userMemoryRepo "go-adr-clean/internal/user/repository/memory"
	userUsecase "go-adr-clean/internal/user/usecase"
	"log"
	"net/http"
)

type Dependency struct {
	Users map[int]userModel.User
}

type Delivery struct {
	Mux *http.ServeMux
}

var (
	userUsecaseSingleton userUsecase.UserUsecase
)

func initInMemoryData() (
	userData map[int]userModel.User,
) {
	userData = make(map[int]userModel.User)
	userData[1] = userModel.User{
		Name: "adr",
		ID:   1,
		Age:  22,
	}

	return
}

func initDependency() Dependency {
	userData := initInMemoryData()
	return Dependency{
		Users: userData,
	}
}

func initDelivery() Delivery {
	return Delivery{
		Mux: http.NewServeMux(),
	}
}

func initUserUsecase(dep Dependency, del Delivery) {
	userMemoryRepo := userMemoryRepo.NewUserMemoryRepo(dep.Users)
	userUsecaseSingleton = userUsecase.NewUserUsecase(userMemoryRepo)
	userDeliveryHttp.NewUserHttpHandler(del.Mux, userUsecaseSingleton)
}

func main() {
	log.Println("starting")
	dep := initDependency()
	del := initDelivery()
	initUserUsecase(dep, del)

	log.Println("server up and running")
	if err := http.ListenAndServe("127.0.0.1:5555", del.Mux); err != nil {
		log.Fatal(err)
	}
}
