package domain

type Karyawan struct {
	Nik        *string `gorm:"column:nik"`
	Nama       *string `gorm:"column:nama"`
	Alamat     *string `gorm:"column:alamat"`
	NoTelp     *string `gorm:"column:no_telp"`
	Email      *string `gorm:"column:email"`
	Foto       *string `gorm:"column:foto"`
	KodeLokasi string  `gorm:"column:kode_lokasi"`
	IdUser     uint    `gorm:"column:id_user"`
}
