package repository

import (
	"context"
	"database/sql"
	"errors"
	"golangapi/helper"
	"golangapi/model/domain"
)

// buat struct yang mengikuti kontrak CategoryRepository di file sebelah
type CategoryRepositoryImplementation struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImplementation{}
}

// save/create
func (repository *CategoryRepositoryImplementation) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//query untuk insert data
	SQL := "insert into category(name) values (?)"         //name saja karena autoincrement si id nya
	result, err := tx.ExecContext(ctx, SQL, category.Name) //untuk mengeksekusi
	helper.PanicIfError(err)                               //dibuatkan package sendiri untuk mrngrcek errornya

	id, err := result.LastInsertId() //mengetahui id terakhir yang dikirim
	helper.PanicIfError(err)

	category.Id = int(id)
	return category //kembalikan categorynya
}

// keterangan sama dengan yang atas
// update
func (repository *CategoryRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

// delete
func (repository *CategoryRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) { //tidak ada return value
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

// findby/cari berdasarkan
func (repository *CategoryRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) { //categorynya int
	SQL := "select id,name from category where id = ?" //harus disebutkan apa yang dipilihnya
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	categori := domain.Category{} //mengkosongkan category
	if rows.Next() {
		err := rows.Scan(&categori.Id, categori.Name) //mengambil nilai asli
		helper.PanicIfError(err)

		return categori, nil
	} else {
		return categori, errors.New("dd not found") //memberi pesan bahwa tidak di temukan
	}
}

// cari semua
func (repository *CategoryRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category { //parameter ketiganya dikosongkan kembalinya slice
	SQL := "select id,name from category"
	rows, err := tx.QueryContext(ctx, SQL) //tidak usah di isi parameternya
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category //membuat slice yang menampung category
	for rows.Next() {
		categori := domain.Category{}
		err := rows.Scan(&categori.Id, categori.Name)
		helper.PanicIfError(err)

		categories = append(categories, categori)
	}

	return categories
}
