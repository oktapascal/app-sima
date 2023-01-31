package domain

import "time"

type Area struct {
	IdArea    string     `gorm:"column:id_area"`
	Name      string     `gorm:"column:name"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	IdUser    int        `gorm:"column:id_user"`
}
