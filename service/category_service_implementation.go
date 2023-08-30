package service

import (
	"context"
	"database/sql"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/model/domain"
	"golang-restfulapi-exercise/model/web"
	"golang-restfulapi-exercise/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImplem struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CategoryServiceImplem) FindByIdGet(ctx context.Context, idKategori int64) web.CategoryResponse {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	categoryid, error := service.CategoryRepository.FindById(ctx, tx, idKategori)
	helper.PanicIfError(error)

	return helper.ToCategoryResponse(categoryid)
}

func (service *CategoryServiceImplem) FindAllGet(ctx context.Context) []web.CategoryResponse {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImplem) CreatePost(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	category := domain.Category{
		Namakategori: request.Namakategori,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImplem) UpdatePut(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	category, error := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(error)

	category.Namakategori = request.Namakategori

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplem) DeleteDelete(ctx context.Context, idKategori int64) {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	category, error := service.CategoryRepository.FindById(ctx, tx, idKategori)
	helper.PanicIfError(error)

	service.CategoryRepository.Delete(ctx, tx, category)
}
