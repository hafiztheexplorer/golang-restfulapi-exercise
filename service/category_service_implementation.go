package service

import (
	"context"
	"database/sql"
	"golang-restfulapi-exercise/exception"
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

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImplem{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImplem) FindById(ctx context.Context, idKategori int) web.CategoryResponse {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	categoryid, error := service.CategoryRepository.FindById(ctx, tx, idKategori)
	if error != nil {
		panic(exception.NewNotFoundError(error.Error()))
	}

	return helper.ToCategoryResponse(categoryid)
}

func (service *CategoryServiceImplem) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImplem) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	error := service.Validate.Struct(request)
	helper.PanicIfError(error)

	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	category := domain.Category{
		Namakategori: request.Namakategori,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImplem) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	error := service.Validate.Struct(request)
	helper.PanicIfError(error)

	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	category, error := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if error != nil {
		panic(exception.NewNotFoundError(error.Error()))
	}

	category.Namakategori = request.Namakategori

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImplem) Delete(ctx context.Context, idKategori int) {
	tx, error := service.DB.Begin()
	helper.PanicIfError(error)
	defer helper.CommitorRollback(tx)

	category, error := service.CategoryRepository.FindById(ctx, tx, idKategori)
	if error != nil {
		panic(exception.NewNotFoundError(error.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}
