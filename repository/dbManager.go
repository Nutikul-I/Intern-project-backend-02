package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

// Replace with your own connection parameters
var server = "54.251.71.182"
var port = 3306
var user = "dev"
var password = "vuBEpYW6Q4iGqm9i"
var database = "dev"
var DB *sql.DB

func Init() {
	server := "54.251.71.182"
	port := 3306
	user := "dev"
	password := "vuBEpYW6Q4iGqm9i"
	database := "dev"
	// Create connection string
	var err error

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
	DB, err = sql.Open("mysql", connString)

	if err != nil {
		log.Error("**** Error creating connection pool: " + err.Error())
	}
	log.Debug("==-- Connected! --==")

}

func ConnectDB() *sql.DB {
	var err error

	if DB == nil {
		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
		DB, err = sql.Open("mysql", connString)

		if err != nil {
			log.Error("**** Error creating connection pool: " + err.Error())
		}
	}

	log.Debug("==-- Connected! --==")
	return DB
}
