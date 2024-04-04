package dao

import "github.com/lib/pq"

type NewCar struct {
	RegNums pq.StringArray `json:"regNums" gorm:"type:text[]"`
	Mark    string         `json:"mark"`
	Model   string         `json:"model"`
	Year    int            `json:"year"`
	Owner   struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
	} `json:"owner"`
}
