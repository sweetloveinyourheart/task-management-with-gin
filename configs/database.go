package configs

import (
	"fmt"
	"os"
	"task-management-with-gin/helpers"
	"task-management-with-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresConnection() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASS")
	dbName := os.Getenv("POSTGRES_DB")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helpers.ErrorPanic(err)

	return db
}

func DatabaseMigration(db *gorm.DB) {
	// Migrate the schema
	err := db.AutoMigrate(&models.User{}, &models.Board{}, &models.Card{}, &models.Comment{}, &models.List{})
	if err != nil {
		helpers.ErrorPanic(err)
	}
}
