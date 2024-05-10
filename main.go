package main

import (
	"NotificationService-rl/dataBase/postgres"
	"NotificationService-rl/gateway"
	"NotificationService-rl/rateLimited"
	"log"
)

func main() {
	dbClient, dbErr := postgres.NewPostgres()
	if dbErr != nil {
		log.Panic(dbErr)
	}
	dbRepository := postgres.NewDbPRepository(dbClient)
	repositoryImp := rateLimited.NewRepository(dbRepository)
	service := rateLimited.NewNotificationServiceImpl(gateway.Gateway{}, repositoryImp)
	service.Send("news", "user", "news 1")
	service.Send("news", "user", "news 2")
	service.Send("news", "user", "news 3")
	service.Send("news", "another user", "news 1")
	service.Send("update", "user", "update 1")
}
