package main

import (
	"github.com/micro/go-micro"
	"github.com/phuonghau98/walkie-user-service/proto"
	"log"
)
type Api struct {
	UserClient user.UserServiceClient
}

func main() {
	service := micro.NewService(
		micro.Name("api.gateway"),
	)

	// parse command line flags
	service.Init()

	_ = service.Server().Handle(
		service.Server().NewHandler(
			&Api{UserClient: user.NewUserServiceClient("service.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
