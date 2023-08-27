package service

import (
	"context"
	"golangapi/model/web"
)

type CategoryService interface { //dibuat kontrak interface dulu
	//mengikuti jumlah APInya
	Create(ctx context.Context, request web.CategoryCreate) web.CategoryResponse //pertama di isi context seperti biasa //kedua yaitu dengan memasukan model, modelnya di model/web (bisa request bisa response)
	Update(ctx context.Context, request web.CategoryUpdate) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
