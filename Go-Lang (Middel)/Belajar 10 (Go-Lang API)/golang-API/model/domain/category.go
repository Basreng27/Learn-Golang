package domain

//Setiap tabel di database harus di presentasikan menjadi struct
type Category struct { //tabel category di database belajar_golang_api
	Id   int
	Name string
}
