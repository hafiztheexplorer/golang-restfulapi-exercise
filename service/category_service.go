package service

import (
	"context"
	"golang-restfulapi-exercise/model/web"
)

type CategoryService interface {
	//Findbyid category spesifik ke id korelasi dengan API GetById
	FindById(ctx context.Context, idKategori int) web.CategoryResponse
	//FindAll category tidak spesifik, korelasi dengan API Get
	FindAll(ctx context.Context) []web.CategoryResponse
	//Create category korelasi dengan API Post
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	//Update/replace category korelasi dengan API Put
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	//Delete category korelasi dengan API Get
	Delete(ctx context.Context, idKategori int)
}
