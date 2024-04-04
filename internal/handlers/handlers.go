package handlers

import (
	"effectivemobile-test/internal/dao"
	"effectivemobile-test/internal/database"
	"effectivemobile-test/internal/models"
	"effectivemobile-test/logs"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

// @Summary Добавление новой машины
// @Description Добавление новой машины в каталог
// @Tags cars
// @Accept json
// @Produce json
// @Param newCar body dao.NewCar true "Данные о новой машине"
// @Success 201 "Машина успешно добавлена"
// @Failure 400 "Некорректный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /addCar [post]
func AddCar(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	logger := logs.GetLogger()

	var newCar dao.NewCar
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var owner models.People
	if err := db.Where(&models.People{Name: newCar.Owner.Name, Surname: newCar.Owner.Surname,
		Patronymic: newCar.Owner.Patronymic}).FirstOrCreate(&owner).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	car := models.Car{
		RegNums: newCar.RegNums,
		Mark:    newCar.Mark,
		Model:   newCar.Model,
		Year:    newCar.Year,
		OwnerID: owner.ID,
	}
	if err := db.Create(&car).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Добавлен новый автомобиль")

	w.WriteHeader(http.StatusCreated)
}

// @Summary Получение списка автомобилей с возможностью пагинации
// @Description Получение списка автомобилей с возможностью пагинации
// @Tags cars
// @Produce json
// @Param page query int false "Номер страницы (по умолчанию 1)"
// @Param limit query int false "Максимальное количество записей на странице (по умолчанию 10)"
// @Success 200 "Успешный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /cars [get]
func GetFilteredCars(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	logger := logs.GetLogger()

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}

	var cars []models.Car
	offset := (pageInt - 1) * limitInt
	db.Offset(offset).Limit(limitInt).Find(&cars)

	w.Header().Set("Content-Type", "application/json")

	logger.Debug("Получен запрос на получение всех автомобилей")

	json.NewEncoder(w).Encode(cars)
}

// @Summary Обновление данных о машине по идентификатору
// @Description Обновление данных о машине по идентификатору
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор машины"
// @Param updatedCar body models.Car true "Обновленные данные о машине"
// @Success 200 "Данные о машине успешно обновлены"
// @Failure 400 "Некорректный запрос"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /update/{id} [put]
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	logger := logs.GetLogger()

	vars := mux.Vars(r)
	id := vars["id"]

	var updatedCar models.Car
	json.NewDecoder(r.Body).Decode(&updatedCar)
	if err := db.Model(&models.Car{}).Where("id = ?", id).Updates(updatedCar).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Данные о машине обновлены")

	w.WriteHeader(http.StatusOK)
}

// @Summary Удаление записи о машине по идентификатору
// @Description Удаление записи о машине по идентификатору
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор машины"
// @Success 200 "Запись успешно удалена"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /delete/{id} [delete]
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	logger := logs.GetLogger()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := db.Where("id = ?", id).Delete(&models.Car{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Запись удалена")

	w.WriteHeader(http.StatusOK)
}
