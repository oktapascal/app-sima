package domain

type Menu struct {
	KodeKlpMenu  string `gorm:"column:kode_klp_menu"`
	NameKlpMenu  string `gorm:"colum:name_klp_menu"`
	ListMenuJson []byte `gorm:"column:menu"`
	ListMenu     interface{}
}
