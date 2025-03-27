package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DatabaseSetting struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func LoadConfig() (*DatabaseSetting, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	host := os.Getenv("DB1_HOST")
	portStr := os.Getenv("DB1_PORT")
	name := os.Getenv("DB1_NAME")
	user := os.Getenv("DB1_USER")
	password := os.Getenv("DB1_PASSWORD")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 5432
	}

	dbConfig := &DatabaseSetting{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: password,
	}

	return dbConfig, nil
}
