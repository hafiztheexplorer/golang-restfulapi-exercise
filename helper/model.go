package helper

import (
	"golang-restfulapi-exercise/model/domain"
	"golang-restfulapi-exercise/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:           category.Id,
		Namakategori: category.Namakategori,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var CategoryResponses []web.CategoryResponse
	for _, category := range categories {
		CategoryResponses = append(CategoryResponses, ToCategoryResponse(category))
	}

	return CategoryResponses
}
