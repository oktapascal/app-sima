package domain

import "time"

type Fm struct {
	IdFm      string     `gorm:"column:id_fm"`
	Name      string     `gorm:"column:name"`
	IdArea    string     `gorm:"column:id_area"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	IdUser    int        `gorm:"column:id_user"`
}
