package domain

type User struct {
	IdUser       int     `gorm:"column:id_user"`
	Username     string  `gorm:"column:username"`
	Password     string  `gorm:"column:password"`
	Role         string  `gorm:"column:role"`
	RedirectView string  `gorm:"column:redirect_view"`
	FlagInput    int     `gorm:"column:flag_input"`
	FlagEdit     int     `gorm:"column:flag_edit"`
	FlagDelete   int     `gorm:"column:flag_delete"`
	Photo        *string `gorm:"column:photo"`
	Provider     *string `gorm:"column:provider"`
	IDProvider   *int    `gorm:"column:id_provider"`
}
