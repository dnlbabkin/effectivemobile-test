package database

import (
	"effectivemobile-test/internal/models"
	"effectivemobile-test/logs"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func migrate(db *gorm.DB) {
	logger := logs.GetLogger()

	db.AutoMigrate(&models.People{}, &models.Car{})
	db.Model(&models.Car{}).AddForeignKey("owner_id", "peoples(id)", "CASCADE", "CASCADE")

	logger.Info("Миграции успешны")
}

func ConnectDB() (*gorm.DB, error) {
	logger := logs.GetLogger()

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	migrate(db)

	logger.Info("Успешное соединение")

	return db, nil
}
