package task

import "time"

type Task struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Task      string     `json:"task"`
	IsDone    bool       `json:"is_done"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
