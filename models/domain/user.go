package domain

type User struct {
	Username   string  `firestore:"username"`
	Password   string  `firestore:"password"`
	Role       string  `firestore:"role"`
	KodeLokasi string  `firestore:"kode_lokasi"`
	Nik        string  `firestore:"nik"`
	Nama       string  `firestore:"nama"`
	Alamat     *string `firestore:"alamat"`
	NoTelp     *string `firestore:"no_telp"`
	Email      *string `firestore:"email"`
	Jabatan    *string `firestore:"jabatan"`
	Foto       *string `firestore:"foto"`
}
