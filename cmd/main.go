package main

import (
	"dota2scheduler/config"
	"dota2scheduler/intenal/handler"
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/server"
	"dota2scheduler/intenal/service"
	"dota2scheduler/redis"
	"log"
	"os"
)

func main() {

	// setting config file for the project
	var configPath string

	switch len(os.Args[1:]) {
	case 1:
		configPath = os.Args[1]
	default:
		configPath = "./.env"
	}

	// setting up log file
	logfile, err := os.Create("app.log")
	if err != nil {
		log.Fatalf("Error while creating log file %s", err)
	}
	defer logfile.Close()

	flags := log.Ldate | log.Ltime | log.Lshortfile

	log.SetOutput(logfile)
	log.SetFlags(flags)

	infoLogger := log.New(logfile, "INFO: ", flags)
	errorLogger := log.New(logfile, "ERROR: ", flags)

	// loading config
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		errorLogger.Fatalf("Error while loading config, %s", err)

	}

	// loading db
	db, err := repository.NewDB(conf)
	if err != nil {
		errorLogger.Fatalf("Error while opening db, %s", err)
	}

	// loading redis
	client, err := redis.NewRedisClient()
	if err != nil {
		errorLogger.Fatalf("Error while opening client %s", err)
	}

	// repo initialization
	repo := repository.NewRepo(db)

	// service initialization
	serv := service.NewService(*repo, client)

	// handler initialization
	controller := handler.NewHandler(serv, infoLogger, errorLogger)

	// initializing scheduler
	go controller.Scheduler()

	errorLogger.Fatal(server.Runserver(conf, controller.Routes(), infoLogger))
}
