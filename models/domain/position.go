package domain

import "time"

type Position struct {
	IdPosition string     `gorm:"column:id_position"`
	Name       string     `gorm:"column:name"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}
