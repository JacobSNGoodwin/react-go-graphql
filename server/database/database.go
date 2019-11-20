package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Database holds connection string info for that database
type Database struct {
	Host    string
	Port    int
	User    string
	Name    string
	SSLMode string
	DB      *gorm.DB
}

// Connect returns the DB connection for the connection info in the struct
func (d *Database) Connect() {
	connStr := "host=%s port=%d user=%s dbname=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, d.Host, d.Port, d.User, d.Name, d.SSLMode)

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=user dbname=gql_demo password=password sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to create connection to postgres database: %v", err.Error())
	}

	d.DB = db
}
