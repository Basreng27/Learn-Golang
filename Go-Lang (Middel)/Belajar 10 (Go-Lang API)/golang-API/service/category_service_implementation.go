package service

import (
	"context"
	"database/sql"
	"golangapi/helper"
	"golangapi/model/domain"
	"golangapi/model/web"
	"golangapi/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImplementation struct {
	CategoryRepository repository.CategoryRepository //membutuhkan repository
	DB                 *sql.DB                       //koneksi kedatabasenya
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, Validate *validator.Validate) CategoryService {
	return &CategoryServiceImplementation{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           Validate}
}

func (service *CategoryServiceImplementation) Create(ctx context.Context, request web.CategoryCreate) web.CategoryResponse {
	err := service.Validate.Struct(request) //sebelum mulai di lakukan validasi
	helper.PanicIfError(err)

	tx, err := service.DB.Begin() //transaksi di mulai ke database
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx) //menggunakan package cek di helper/tx.go

	category := domain.Category{ //domainnya
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplementation) Update(ctx context.Context, request web.CategoryUpdate) web.CategoryResponse {
	err := service.Validate.Struct(request) //sebelum mulai di lakukan validasi
	helper.PanicIfError(err)

	tx, err := service.DB.Begin() //transaksi di mulai ke database
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx) //menggunakan package cek di helper/tx.go

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err) //mengecek ada atau tidak adany adata

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplementation) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin() //transaksi di mulai ke database
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx) //menggunakan package cek di helper/tx.go

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err) //mengecek ada atau tidak adany adata

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImplementation) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin() //transaksi di mulai ke database
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx) //menggunakan package cek di helper/tx.go

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplementation) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin() //transaksi di mulai ke database
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx) //menggunakan package cek di helper/tx.go}

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories) //folder helper/model
}
