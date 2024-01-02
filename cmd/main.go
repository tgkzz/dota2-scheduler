package main

import (
	"dota2scheduler/config"
	"dota2scheduler/intenal/handler"
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/server"
	"dota2scheduler/intenal/service"
	"log"
)

// TODO: fill layer have methods that is logically must not be their, make it in another package

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error while loading config, %s", err)
	}

	db, err := repository.NewDB(conf)
	if err != nil {
		log.Fatalf("Error while opening db, %s", err)
	}

	// repo initialization
	repo := repository.NewRepo(db)

	// service initialization
	serv := service.NewService(*repo)

	// handler initialization
	controller := handler.NewHandler(serv)

	// initializing scheduler
	go controller.Scheduler()

	log.Fatal(server.Runserver(conf, controller.Routes()))
}
