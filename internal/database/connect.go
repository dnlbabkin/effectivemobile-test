package database

import (
	"effectivemobile-test/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.People{}, &models.Car{})
	db.Model(&models.Car{}).AddForeignKey("owner_id", "peoples(id)", "CASCADE", "CASCADE")
}

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	migrate(db)

	log.Println("Успешное соединение")

	return db, nil
}
