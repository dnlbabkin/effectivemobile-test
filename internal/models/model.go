package models

import "github.com/lib/pq"

type Car struct {
	ID      uint           `gorm:"primary_key" json:"id"`
	RegNums pq.StringArray `gorm:"type:text[]" json:"regNums"`
	Mark    string         `gorm:"type:varchar(255)" json:"mark"`
	Model   string         `gorm:"type:varchar(255)" json:"model"`
	Year    int            `json:"year"`
	OwnerID uint           `json:"owner_id"`
	Owner   People         `json:"owner"`
}

type People struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Name       string `gorm:"type:varchar(100)" json:"name"`
	Surname    string `gorm:"type:varchar(100)" json:"surname"`
	Patronymic string `gorm:"type:varchar(100)" json:"patronymic"`
	Cars       []Car  `json:"cars"`
}
