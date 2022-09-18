package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Psql_config struct {
	PORT     string
	HOST     string
	PASSWORD string
}

func getDatabaseString(host, port, user, password, database, sslMode string) string {

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, database, sslMode)
}

func NewPsqlConfig(host, port, user, password, database, sslMode string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(getDatabaseString(host, port, user, password, database, sslMode)))

	if err != nil {
		panic("SQL Connection error")
	}

	return db
}
