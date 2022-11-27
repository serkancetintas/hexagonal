package main

import (
	"hexagonal/internals/core/services"
	"hexagonal/internals/handlers"
	"hexagonal/internals/repositories"
	"hexagonal/internals/server"
)

func main() {
	mongoConn := "secret🤫"
	//repositories
	userRepository, _ := repositories.NewUserRepository(mongoConn)
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	//server
	newServer := server.NewServer(
		userHandlers,
	)
	newServer.Initialize()
}
