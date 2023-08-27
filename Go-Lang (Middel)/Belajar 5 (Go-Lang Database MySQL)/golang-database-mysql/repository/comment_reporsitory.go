package repository

import (
	"context"
	"golang-database-mysql/entity"
)

type CommentRepository interface { //untuk nama repository biasanya sama dengan nama entity untuk lebih mudah
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error )
}
