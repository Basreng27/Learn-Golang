package repository

import (
	"context"
	"database/sql"
	"golangapi/model/domain"
)

// kontrak dulu dengan interface //ontrak dengan semua function yang dibutuhkan
type CategoryRepository interface { //buat setiap function untuk api
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category    //untuk create //untuk parameternya biasakan menggunakan context //dengan variabel ctx //parameter selanjutnya biasanya harus tansaksional(nanti di seeting transaksional atau tidaknya/readonly atau tidaknya)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category  //untuk update //sql.TX = Transaksional //harus pointer karena struct jika tidak maka akn tercopt bukan data asli //dan parameter ketiga adalah data yang aslinya
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)                  //untuk delete //untuk return valuenya
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) //untuk mencari dari id //karna hanya mencari data
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category                         //untuk mencari semua data //return slice mencari semua
}
