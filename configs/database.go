package configs

import (
	"fmt"
	"os"
	"task-management-with-gin/helpers"
	"task-management-with-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitPostgresConnection() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASS")
	dbName := os.Getenv("POSTGRES_DB")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		helpers.ErrorPanic(err)
	}

	database = db
	return db
}

func MigrateDatabase(db *gorm.DB) {
	// Migrate the schema
	err := db.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.Card{},
		&models.Comment{},
		&models.List{})

	if err != nil {
		helpers.ErrorPanic(err)
	}
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return database
}
