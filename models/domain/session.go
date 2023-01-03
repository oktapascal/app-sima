package domain

import "time"

type Session struct {
	AuthToken string     `gorm:"column:auth_token"`
	IdUser    uint       `gorm:"column:id_user"`
	CreatedAt time.Time  `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
