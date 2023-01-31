package domain

import "time"

type Bm struct {
	IdBm      string     `gorm:"column:id_bm"`
	Name      string     `gorm:"column:name"`
	IdFm      string     `gorm:"column:id_fm"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	IdUser    int        `gorm:"column:id_user"`
}
