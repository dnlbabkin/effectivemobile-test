package handlers

import (
	"effectivemobile-test/internal/dao"
	"effectivemobile-test/internal/database"
	"effectivemobile-test/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// @Summary Добавление новой машины
// @Description Добавление новой машины в каталог
// @Tags cars
// @Accept json
// @Produce json
// @Param car body Car true "Название машины и владельца"
// @Success 201
// @Router /addCar [post]
func AddCar(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

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

	w.WriteHeader(http.StatusCreated)
}

// @Summary Возвращение отфильтрованного списка машин
// @Description Вазваращает данные с фильтрацией и пагинацией
// @Tags cars
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы для пагинации"
// @Param limit query int false "Количество элементов на странице"
// @Success 200 {array} Car
// @Router /cars [get]
func GetFilteredCars(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

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

	json.NewEncoder(w).Encode(cars)
}

// @Summary Изменение данных по id
// @Description Изменение одного или нескольких полей по id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "id машины"
// @Param car body Car true "Поле для обновления"
// @Success 200
// @Router /update/{id} [put]
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var updatedCar models.Car
	json.NewDecoder(r.Body).Decode(&updatedCar)
	if err := db.Model(&models.Car{}).Where("id = ?", id).Updates(updatedCar).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Удаление машины по id
// @Description Удалить машину по id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "id машины"
// @Success 200
// @Router /delete/{id} [delete]
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := db.Where("id = ?", id).Delete(&models.Car{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
