package web

type CategoryCreate struct {
	Name string `validate:"required,max=255,min=1" json:"name"` //karena idnya autoincrement //validasinya harus di wajibkan di isi
}
