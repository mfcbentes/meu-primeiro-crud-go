package main

import (
	"github.com/mfcbentes/meu-primeiro-crud-go/src/controller"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model/repository"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
