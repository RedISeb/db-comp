package db

import (
	"database/sql"
	"fmt"
	"github.com/rediseb/db-comp/config"
	"log"
)

func Connect(config *config.DatabaseSetting) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.User, config.Password, config.Name, config.Host, config.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Successfully connected to the database!")
	return db, nil
}
