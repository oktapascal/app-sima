package domain

import "time"

type Session struct {
	IdSession int        `gorm:"column:id_session"`
	AuthToken string     `gorm:"column:auth_token"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	IdUser    int        `gorm:"column:id_user"`
}
