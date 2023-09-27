package main

import (
	"log"

	"github.com/Rest_Savings/apps/models"
	"github.com/Rest_Savings/pkg/database"
)

func main() {
	db, err := database.New().GetInstance()
	if err != nil {
		log.Print(err)
	}
	db.AutoMigrate(
		&models.Account{},
		&models.Trasaction{},
	)
	conn, _ := db.DB()
	defer conn.Close()
}
