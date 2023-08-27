package helper

import (
	"golangapi/model/domain"
	"golangapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoyResponses []web.CategoryResponse
	for _, category := range categories {
		categoyResponses = append(categoyResponses, ToCategoryResponse(category)) //dikonversi
	}

	return categoyResponses
}
