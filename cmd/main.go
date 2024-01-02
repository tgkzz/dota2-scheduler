package main

import (
	"dota2scheduler/config"
	"dota2scheduler/intenal/handler"
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/server"
	"dota2scheduler/intenal/service"
	"log"
	"os"
)

// TODO: fill layer have methods that is logically must not be their, make it in another package

func main() {
	var configPath string

	switch len(os.Args[1:]) {
	case 1:
		configPath = os.Args[1]
	default:
		configPath = "./.env"
	}

	conf, err := config.LoadConfig(configPath)
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
