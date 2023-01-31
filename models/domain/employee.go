package domain

type Employee struct {
	IdEmployee  int     `gorm:"column:id_employee"`
	Nik         string  `gorm:"column:nik"`
	IdLocation  string  `gorm:"column:id_location"`
	Name        string  `gorm:"column:name"`
	Address     *string `gorm:"column:address"`
	NoTelp      *string `gorm:"column:no_telp"`
	Email       *string `gorm:"column:email"`
	ActiveFlags int     `gorm:"column:active_flags"`
	IdArea      *string `gorm:"column:id_area"`
	IdFm        *string `gorm:"column:id_fm"`
	IdBm        *string `gorm:"column:id_bm"`
	IdPosition  *string `gorm:"column:id_position"`
	IdUser      int     `gorm:"column:id_user"`
}
