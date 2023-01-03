package domain

type User struct {
	Id       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Role     string
	Karyawan Karyawan  `gorm:"foreignKey:IdUser;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Session  []Session `gorm:"foreignKey:IdUser;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
