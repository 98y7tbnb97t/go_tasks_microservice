package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Task      string     `json:"task"`
	IsDone    bool       `json:"is_done"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Task{})
}
