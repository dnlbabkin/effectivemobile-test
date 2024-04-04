package main

import (
	_ "effectivemobile-test/docs"
	"effectivemobile-test/internal/handlers"
	"effectivemobile-test/logs"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Каталог новых автомобилей
// @version 1.0
// @description API для каталога новых автомобилей

// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	logger := logs.GetLogger()

	logger.Info("Запуск")
	router := mux.NewRouter()

	router.HandleFunc("/addCar", handlers.AddCar).Methods("POST")
	router.HandleFunc("/cars", handlers.GetFilteredCars).Methods("GET")
	router.HandleFunc("/update/{id}", handlers.UpdateCar).Methods("PUT")
	router.HandleFunc("/delete/{id}", handlers.DeleteCar).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	logger.Info("Сервер прослушивается на порту 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
