package config

import (
	"fmt"
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

func LoadConfig() (map[string]*DatabaseSetting, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	dbConfigs := make(map[string]*DatabaseSetting)

	for i := 1; ; i++ {
		keyPrefix := fmt.Sprintf("DB%d", i)
		host := os.Getenv(fmt.Sprintf("%s_HOST", keyPrefix))
		if host == "" {
			break
		}

		portStr := os.Getenv(fmt.Sprintf("%s_PORT", keyPrefix))
		name := os.Getenv(fmt.Sprintf("%s_NAME", keyPrefix))
		user := os.Getenv(fmt.Sprintf("%s_USER", keyPrefix))
		password := os.Getenv(fmt.Sprintf("%s_PASSWORD", keyPrefix))

		port, err := strconv.Atoi(portStr)
		if err != nil {
			port = 5432
		}

		dbConfigs[keyPrefix] = &DatabaseSetting{
			Host:     host,
			Port:     port,
			Name:     name,
			User:     user,
			Password: password,
		}
	}

	return dbConfigs, nil
}
