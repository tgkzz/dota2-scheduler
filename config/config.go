package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host string
	Port string
	DB   DB
}

type DB struct {
	DriverName     string
	DataSourceName string
	SQL            string
}

func LoadConfig(path string) (Config, error) {

	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}
	config := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DB: DB{
			DriverName:     os.Getenv("DRIVERNAME"),
			DataSourceName: os.Getenv("DATASOURCENAME"),
			SQL:            os.Getenv("SQL"),
		},
	}
	fmt.Println(config)

	return config, nil
}
