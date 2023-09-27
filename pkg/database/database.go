package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type GormDB struct {
	Db   *gorm.DB
	Conn *sql.DB
}

func New() *GormDB {
	return &GormDB{}
}

func (godb *GormDB) GetInstance() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	destination := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_Port"),
		os.Getenv("DB_User"),
		os.Getenv("DB_Password"),
		os.Getenv("DB_Name"),
		os.Getenv("DB_TimeZone"),
	)

	lock.Lock()
	godb.Db, err = gorm.Open(postgres.Open(destination), &gorm.Config{})
	lock.Unlock()
	if err != nil {
		log.Printf("Error %s when connection DB\n", err)
		log.Fatal(err)

	}
	log.Printf("DB Connect")
	return godb.Db, nil
}
