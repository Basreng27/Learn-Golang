package repository

import "golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category //function bernama "FindById". nilai balikan kalau ada akan menggunakan "Category" jika tidak ada akan menggunakan "*entity" akan nil hasilnya
}
