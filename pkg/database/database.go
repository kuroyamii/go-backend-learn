package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase(username string, password string, address string, dbName string) *sql.DB {
	log.Println("INFO GetDatabase database connection: starting database connection process")
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		log.Printf("ERROR GetDatabase sql open connection fatal error: %v", err)
		for {
			log.Println("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			db, err = sql.Open("mysql", dbConn)
			if err == nil {
				break
			}
		}
	}

	if err = db.Ping(); err != nil {
		log.Printf("ERROR GetDatabase db ping fatal error: %v", err)
		for {
			log.Println("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			db, err = sql.Open("mysql", dbConn)
			err2 := db.Ping()
			if err == nil && err2 == nil {
				break
			}
		}
	}
	log.Printf("INFO GetDatabase database connection: established successfully with %s\n", dbConn)
	return db
}
